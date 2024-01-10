// Manipulate bytes.
package bytes

var (
	Digits      = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	LetterLower = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	LetterUpper = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
)

// isBetween returns whether c is between min and max (inclusive).
func isBetween(c, min, max byte) bool {
	return c >= min && c <= max
}

// IsDigit returns whether c is a number character ascii code.
func IsDigit(c byte) bool {

	return isBetween(c, '0', '9')
}

// IsHexa returns whether c is a hexadecimal (0-9 a-f A-F) character ascii code.
func IsHexa(c byte) bool {
	return IsDigit(c) || isBetween(c, 'a', 'f') || isBetween(c, 'A', 'F')
}

// IsLowerLetter returns whether c is a lower case letter (a-z) character ascii code.
func IsLowerLetter(c byte) bool {
	return isBetween(c, 'a', 'z')
}

// IsUpperLetter returns whether c is an upper case letter (A-Z)  character ascii code.
func IsUpperLetter(c byte) bool {
	return isBetween(c, 'A', 'Z')
}
