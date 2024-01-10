package portscan

// type Config struct {
// 	Target  net.IP         // Target IP address.
// 	Port    uint32         // Target port.
// 	SrcPort uint32         // Source port to use. If 0,
// 	Iface   *net.Interface // Interface to use. If left nil, the select one automatically.
// 	Timeout time.Duration  //
// 	sfd     int            // File descriptor to send packet
// }

// func SendPacket(c Config) error {

// 	var (
// 		fd  int
// 		err error
// 	)

// 	switch {
// 	case ip.IsValid4(c.Target):
// 		fd, err = unix.Socket(unix.AF_INET, unix.SOCK_RAW, unix.IPPROTO_RAW)
// 	case ip.IsValid6(c.Target):
// 		fd, err = unix.Socket(unix.AF_INET6, unix.SOCK_RAW, unix.IPPROTO_RAW)
// 	default:
// 		return fmt.Errorf("invalid Target: %s", c.Target)
// 	}

// 	if err != nil {

// 	}

// 	return nil
// }

// func Scan(c Config) (context.CancelFunc, <-chan gopacket.Packet, error) {

// }
