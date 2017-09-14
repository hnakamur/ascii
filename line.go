package asciibytes

// GetLine returns a line including the last '\n' byte.
func GetLine(buf []byte) []byte {
	for i, b := range buf {
		if b == '\n' {
			return buf[:i+1]
		}
	}
	return buf
}
