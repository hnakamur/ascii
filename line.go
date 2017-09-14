package asciibytes

import "bytes"

// GetLine returns a line including the last '\n' byte.
func GetLine(buf []byte) []byte {
	i := bytes.IndexByte(buf, '\n')
	if i == -1 {
		return buf
	}
	return buf[:i+1]
}
