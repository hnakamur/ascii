package asciibytes

// IsSpace returns true if b is one of '\t', '\n', '\v', '\f', '\r', or ' '.
// It returns false otherwise.
func IsSpace(b byte) bool {
	switch b {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	default:
		return false
	}
}
