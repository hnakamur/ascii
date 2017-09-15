package ascii

// GetLine searches the first '\n' and returns a line including '\n'.
func GetLine(buf []byte) []byte {
	for i, b := range buf {
		if b == '\n' {
			return buf[:i+1]
		}
	}
	return buf
}
