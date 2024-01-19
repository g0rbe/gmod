package freax

import "bytes"

const (
	PathMax = 4096 // chars in a path name including nul
)

// PathJoin joins any number of path elements into a single path, separating them with slashes.
// If the elem is empty or all its elements are empty, returns an empty string.
func PathJoin(elem ...string) string {

	p := bytes.NewBuffer(make([]byte, 0, PathMax))

	for i := range elem {

		// Skip empty elements
		if len(elem[i]) == 0 {
			continue
		}

		p.WriteString(elem[i])

		// Dont add '/' if this is the last elem
		if i == len(elem)-1 {
			break
		}

		// Add '/' to the elem if this elem is not the last
		if elem[i][len(elem[i])-1] != '/' {
			p.WriteByte('/')
		}

	}

	return p.String()
}
