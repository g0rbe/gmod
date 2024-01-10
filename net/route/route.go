package route

// flags from net/route.h

// IPv4 route flags
const (
	RTF_UP         = 0x0001  // Route usable.
	RTF_GATEWAY    = 0x0002  // Destination is a gateway.
	RTF_HOST       = 0x0004  // Host entry (net otherwise).
	RTF_REINSTATE  = 0x0008  // Reinstate route after timeout.
	RTF_DYNAMIC    = 0x0010  // Created dyn. (by redirect).
	RTF_MODIFIED   = 0x0020  // Modified dyn. (by redirect).
	RTF_MTU        = 0x0040  // Specific MTU for this route.
	RTF_MSS        = RTF_MTU // Compatibility.
	RTF_WINDOW     = 0x0080  // Per route window clamping.
	RTF_IRTT       = 0x0100  // Initial round trip time.
	RTF_REJECT     = 0x0200  // Reject route.
	RTF_STATIC     = 0x0400  // Manually injected route.
	RTF_XRESOLVE   = 0x0800  // External resolver.
	RTF_NOFORWARD  = 0x1000  // Forwarding inhibited.
	RTF_THROW      = 0x2000  // Go to next class.
	RTF_NOPMTUDISC = 0x4000  // Do not send packets with DF.
)

// IPv6 route flags
const (
	RTF_DEFAULT     = 0x00010000 // default - learned via ND
	RTF_ALLONLINK   = 0x00020000 // fallback, no routers on link
	RTF_ADDRCONF    = 0x00040000 // addrconf route - RA
	RTF_LINKRT      = 0x00100000 // link specific - device match
	RTF_NONEXTHOP   = 0x00200000 // route with no nexthop
	RTF_CACHE       = 0x01000000 // cache entry
	RTF_FLOW        = 0x02000000 // flow significant route
	RTF_POLICY      = 0x04000000 // policy route
	RTCF_VALVE      = 0x00200000
	RTCF_MASQ       = 0x00400000
	RTCF_NAT        = 0x00800000
	RTCF_DOREDIRECT = 0x01000000
	RTCF_LOG        = 0x02000000
	RTCF_DIRECTSRC  = 0x04000000
	RTF_LOCAL       = 0x80000000
	RTF_INTERFACE   = 0x40000000
	RTF_MULTICAST   = 0x20000000
	RTF_BROADCAST   = 0x10000000
	RTF_NAT         = 0x08000000
)
