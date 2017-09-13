package asciibytes

var asciiSpace = [256]bool{'\t': true, '\n': true, '\v': true, '\f': true, '\r': true, ' ': true}

// IsSpace returns true if b is one of '\t', '\n', '\v', '\f', '\r', or ' '.
// It returns false otherwise.
func IsSpace(b byte) bool {
	return asciiSpace[b]
}
