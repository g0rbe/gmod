package arp

import (
	"net"
)

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

// ifaceNameToBytes converts net.Interface.Name to [16]byte for SIOCSARP.
func ifaceNameToBytes(name string) [16]byte {

	var v [16]byte

	b := []byte(name)

	for i := 0; i < 15 && i < len(b); i++ {
		v[i] = b[i]
	}

	return v
}

// ipv4ToBytes converts ip to [4]byte for SIOCSARP.
// It uses ip.To4() to be sure.
func ipToArpPaAddr(ip net.IP) [4]byte {

	var v [4]byte

	ip = ip.To4()

	for i := 0; i < 4 && i < len(ip); i++ {
		v[i] = ip[i]
	}

	return v
}

// hwAddrToInt8 convert net.HardwareAddr to [14]int8 for SIOCSARP.
func hwAddrToArpHaData(mac net.HardwareAddr) [14]int8 {

	var v [14]int8

	for i := 0; i < 14 && i < len(mac); i++ {
		v[i] = int8(mac[i])
	}

	return v
}

// int8ToHwAddr convert [14]int8 to net.HardwareAddr for SIOCGARP.
func arpHaDataToHwAddr(data [14]int8) net.HardwareAddr {

	var v net.HardwareAddr

	for i := 0; i < 6; i++ {

		v = append(v, byte(data[i]))
	}

	return v
}
