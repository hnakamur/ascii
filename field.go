package asciibytes

// NextField first skips spaces and returns the start and the end index of the field in buf.
func NextField(buf []byte) (start, end int) {
	i := 0
	for IsSpace(buf[i]) {
		i++
	}
	j := i + 1
	for !IsSpace(buf[j]) {
		j++
	}
	return i, j
}
