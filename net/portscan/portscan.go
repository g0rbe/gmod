package portscan

type State uint8

const (
	Open State = 1 << iota
	Closed
	Filtered
	OpenFiltered
)

func (s State) String() string {

	switch s {
	case Open:
		return "Open"
	case Closed:
		return "Closed"
	case Filtered:
		return "Filtered"
	case OpenFiltered:
		return "Open Filtered"
	default:
		return "Unknown"
	}
}
