package arp

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"

	gip "github.com/g0rbe/gmod/net/ip"
	"github.com/g0rbe/gmod/net/route"
)

// See more at "man 7 arp"

type arpreq struct {
	ArpPa   syscall.RawSockaddrInet4
	ArpHa   syscall.RawSockaddr
	Flags   int32
	Netmask syscall.RawSockaddr
	Dev     [16]byte
}

// GetARPCache look for the entry associated to ip and iface from the kernel's cache.
// If iface is nil, select the one automatically.
func GetARPCache(ip net.IP, dev *net.Interface) (Entry, error) {

	var entry Entry

	if ip == nil {
		return entry, fmt.Errorf("ip is nil")
	}
	if !gip.IsValid4(ip) {
		return entry, fmt.Errorf("not valid IPv4")
	}

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return entry, err
	}
	defer syscall.Close(fd)

	if dev == nil {
		route, err := route.GetRoute4(ip)
		if err != nil {
			return entry, fmt.Errorf("failed to get route: %s", err)
		}

		dev = route.Interface
	}

	// Create SIOCGARP request.
	req := arpreq{
		ArpPa: syscall.RawSockaddrInet4{
			Family: syscall.AF_INET,
			Addr:   ipToArpPaAddr(ip),
		},
		Dev: ifaceNameToBytes(dev.Name),
	}

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.SIOCGARP, uintptr(unsafe.Pointer(&req)))

	if errno > 0 {
		if errno == 6 {
			return entry, ErrNotFound
		}
		return entry, errno
	}

	entry.IP = ip
	entry.Type = req.ArpHa.Family
	entry.Flags = req.Flags
	entry.Addr = arpHaDataToHwAddr(req.ArpHa.Data)
	entry.Mask = req.Netmask
	entry.Device = dev

	return entry, nil
}

// SetARPCache set the ARP entry in the kernel's cache.
func SetARPCache(entry Entry) error {

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)

	arpReq := arpreq{
		ArpPa: syscall.RawSockaddrInet4{
			Family: syscall.AF_INET,
			Addr:   ipToArpPaAddr(entry.IP),
		},
		ArpHa: syscall.RawSockaddr{
			Family: uint16(entry.Type),
			Data:   hwAddrToArpHaData(entry.Addr),
		},
		Netmask: entry.Mask,
		Flags:   entry.Flags,
		Dev:     ifaceNameToBytes(entry.Device.Name),
	}

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.SIOCSARP, uintptr(unsafe.Pointer(&arpReq)))

	if errno != 0 {
		return errno
	}

	return nil
}

// DelARPCache removes the ARP entry from the kernel's cache.
func DelARPCache(entry Entry) error {

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)

	arpReq := arpreq{
		ArpPa: syscall.RawSockaddrInet4{
			Family: syscall.AF_INET,
			Addr:   ipToArpPaAddr(entry.IP),
		},
		ArpHa: syscall.RawSockaddr{
			Family: uint16(entry.Type),
			Data:   hwAddrToArpHaData(entry.Addr),
		},
		Netmask: entry.Mask,
		Flags:   entry.Flags,
		Dev:     ifaceNameToBytes(entry.Device.Name),
	}

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.SIOCDARP, uintptr(unsafe.Pointer(&arpReq)))

	if errno != 0 {
		return errno
	}

	return nil
}
