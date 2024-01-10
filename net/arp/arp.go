package arp

import (
	"errors"
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/g0rbe/gmod/net/route"
)

var (
	// Requested IP is not is not in cache.
	ErrNotFound = errors.New("no such device or address")
)

// ARP Flags values. See net/if_arp.h.
const (

	// Completed entry (ha valid).
	ATF_COM int32 = 0x02

	// Permanent entry.
	ATF_PERM int32 = 0x04

	// Publish entry.
	ATF_PUBL int32 = 0x08

	// Has requested trailers.
	ATF_USETRAILERS int32 = 0x10

	// Want to use a netmask (only for proxy entries).
	ATF_NETMASK int32 = 0x20

	// Don't answer this addresses.
	ATF_DONTPUB int32 = 0x40

	// Automatically added entry.
	ATF_MAGIC int32 = 0x80
)

// ARP protocol HARDWARE identifiers.
const (
	ARPHRD_NETROM     uint16 = 0  /* From KA9Q: NET/ROM pseudo. */
	ARPHRD_ETHER      uint16 = 1  /* Ethernet 10/100Mbps.  */
	ARPHRD_EETHER     uint16 = 2  /* Experimental Ethernet.  */
	ARPHRD_AX25       uint16 = 3  /* AX.25 Level 2.  */
	ARPHRD_PRONET     uint16 = 4  /* PROnet token ring.  */
	ARPHRD_CHAOS      uint16 = 5  /* Chaosnet.  */
	ARPHRD_IEEE802    uint16 = 6  /* IEEE 802.2 Ethernet/TR/TB.  */
	ARPHRD_ARCNET     uint16 = 7  /* ARCnet.  */
	ARPHRD_APPLETLK   uint16 = 8  /* APPLEtalk.  */
	ARPHRD_DLCI       uint16 = 15 /* Frame Relay DLCI.  */
	ARPHRD_ATM        uint16 = 19 /* ATM.  */
	ARPHRD_METRICOM   uint16 = 23 /* Metricom STRIP (new IANA id).  */
	ARPHRD_IEEE1394   uint16 = 24 /* IEEE 1394 IPv4 - RFC 2734.  */
	ARPHRD_EUI64      uint16 = 27 /* EUI-64.  */
	ARPHRD_INFINIBAND uint16 = 32 /* InfiniBand.  */

)

// Represent an ARP entry
type Entry struct {
	IP     net.IP              //
	Type   uint16              //
	Flags  int32               //
	Addr   net.HardwareAddr    // MAC
	Mask   syscall.RawSockaddr // Used in Proxy ARP. Not implemented but returned, somebody maybe find it useful.
	Device *net.Interface      // Entry valid for this interface
}

// The string is separated with a \t.
// The Mask is ignored.
func (c Entry) String() string {

	return fmt.Sprintf("%s\t0x%x\t0x%x\t%s\t%s",
		c.IP.String(),
		c.Type,
		c.Flags,
		c.Addr.String(),
		c.Device.Name)
}

// Get query the kernel's cache first.
// If the entry is not in the cache, do an ARP request and registers the entry in the cache (if CAP_NET_ADMIN is set).
func Get(ip net.IP, timeout time.Duration) (Entry, error) {

	rt, err := route.GetRoute4(ip)
	if err != nil {
		return Entry{}, fmt.Errorf("failed to get route: %s", err)
	}

	// cache, err := GetARPCache(ip, rt.Interface)
	// if err != nil {

	// 	// Not in cahce, do a request
	// 	if errors.Is(err, ErrNotFound) {
	return Request(ip, rt.Interface, timeout)
	// 	}

	// 	return Entry{}, err
	// }

	// return cache, err
}
