package arp

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/g0rbe/gmod/net/capability"
	giface "github.com/g0rbe/gmod/net/iface"
	"github.com/g0rbe/gmod/net/route"
	"github.com/google/gopacket"
	"github.com/google/gopacket/afpacket"
	"github.com/google/gopacket/layers"
	"golang.org/x/sys/unix"
)

var (
	BroadcastMAC = net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	ZeroMAC      = net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

type dataSource struct {
	Fd    int
	buff  []byte
	iface *net.Interface
}

// Used in read() as a DacketDataSource
func newDataSource(fd int, iface *net.Interface) *dataSource {
	return &dataSource{Fd: fd, buff: make([]byte, 1<<16), iface: iface}
}

// Implement gopacket.PacketDataSource
func (s *dataSource) ReadPacketData() (data []byte, ci gopacket.CaptureInfo, err error) {

	data = make([]byte, 65536)

	n, _, err := unix.Recvfrom(s.Fd, data, 0)
	if err != nil {
		return data, ci, err
	}

	data = data[:n]

	return
}

func bindToDev(fd int, dev *net.Interface) error {

	// sockaddr_ll
	sll := &unix.SockaddrLinklayer{
		Ifindex:  dev.Index,
		Protocol: htons(unix.ETH_P_ALL),
	}

	return unix.Bind(fd, sll)
}

// Create the ARP packet: Layer2 + Layer3
func createARP(q net.IP, dev *net.Interface) ([]byte, error) {

	eth := layers.Ethernet{
		DstMAC:       BroadcastMAC,
		SrcMAC:       dev.HardwareAddr,
		EthernetType: layers.EthernetTypeARP,
	}

	devNets, err := giface.GetIPNets4(dev)
	if err != nil {
		return nil, fmt.Errorf("failed to get local IP: %s", err)
	}

	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         0x001,
		SourceHwAddress:   dev.HardwareAddr,
		SourceProtAddress: devNets[0].IP.To4(),
		DstHwAddress:      ZeroMAC,
		DstProtAddress:    q.To4(),
	}

	buff := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{FixLengths: true, ComputeChecksums: true}

	err = gopacket.SerializeLayers(buff, opts, &eth, &arp)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func send(ip net.IP, dev *net.Interface) error {

	sock, err := unix.Socket(unix.AF_PACKET, unix.SOCK_RAW, int(htons(unix.ETH_P_ALL)))
	if err != nil {
		return err
	}
	defer unix.Close(sock)

	if err := bindToDev(sock, dev); err != nil {
		return fmt.Errorf("failed to bind %s: %s", dev.Name, err)
	}

	b, err := createARP(ip, dev)
	if err != nil {
		return fmt.Errorf("failed to create ARP: %s", err)
	}

	// sockaddr_ll
	sll := &unix.SockaddrLinklayer{
		Ifindex: dev.Index,
		Halen:   6, // IEEE 802 MAC-48
		Addr:    [8]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	}

	return unix.Sendto(sock, b, 0, sll)
}

func read(ip net.IP, dev *net.Interface, ctx context.Context, addrCh chan<- *net.HardwareAddr, errCh chan<- error) {

	var (
		eth = layers.Ethernet{}
		arp = layers.ARP{}

		parser  = gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &arp)
		decoded = []gopacket.LayerType{}
	)

	s, err := afpacket.NewTPacket()
	if err != nil {
		errCh <- fmt.Errorf("failed to create tpacket: %s", err)
		close(addrCh)
		close(errCh)
		return
	}

	parser.IgnoreUnsupported = true

	source := gopacket.NewPacketSource(s, layers.LayerTypeEthernet)

	for {

		select {
		case <-ctx.Done():
			errCh <- fmt.Errorf("i/o timeout")
			close(addrCh)
			close(errCh)
			return
		case packetData := <-source.Packets():

			if err := parser.DecodeLayers(packetData.Data(), &decoded); err != nil {
				errCh <- fmt.Errorf("could not decode layers: %s\n", err)
				close(errCh)
				close(addrCh)
				return
			}

			for i := range decoded {
				if decoded[i] == layers.LayerTypeARP {

					if ip.To4().Equal(arp.SourceProtAddress) {
						addrCh <- &net.HardwareAddr{
							arp.SourceHwAddress[0], arp.SourceHwAddress[1],
							arp.SourceHwAddress[2], arp.SourceHwAddress[3],
							arp.SourceHwAddress[4], arp.SourceHwAddress[5]}
						close(errCh)
						close(addrCh)
						return
					}
				}
			}
		}
	}
}

// Request do an ARP request with raw socket.
// Returns the MAC associated to ip.
// If dev is nil, automatically select one.
// If CAP_NET_ADMIN is set, this function registers the entry to the kernel's ARP cache.
func Request(ip net.IP, dev *net.Interface, timeout time.Duration) (Entry, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	if dev == nil {
		route, err := route.GetRoute4(ip)
		if err != nil {
			return Entry{}, fmt.Errorf("failed to get gateway: %s", err)
		}

		dev = route.Interface
	}

	addrCh := make(chan *net.HardwareAddr)
	errCh := make(chan error)

	go read(ip, dev, ctx, addrCh, errCh)

	time.Sleep(500 * time.Millisecond)

	if err := send(ip, dev); err != nil {
		return Entry{}, fmt.Errorf("failed to send: %s", err)
	}

	select {
	case err := <-errCh:
		return Entry{}, err
	case addrs := <-addrCh:
		entry := Entry{
			IP:     ip,
			Type:   ARPHRD_ETHER,
			Flags:  ATF_COM,
			Addr:   net.HardwareAddr{(*addrs)[0], (*addrs)[1], (*addrs)[2], (*addrs)[3], (*addrs)[4], (*addrs)[5]},
			Device: dev}

		// Check capability to set ARP cache
		ok, err := capability.CapabilityCheck(capability.CAP_NET_ADMIN, os.Getpid())
		if err != nil {
			return entry, err
		}
		if ok {
			err = SetARPCache(entry)
		}

		return entry, err
	}
}
