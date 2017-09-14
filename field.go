package ascii

// NextField first skips spaces and returns the start and the end index of the field in buf.
func NextField(buf []byte) (start, end int) {
	i := 0
	l := len(buf)
	for i < l && IsSpace(buf[i]) {
		i++
	}
	if i+1 >= len(buf) {
		return i, i
	}

	j := i + 1
	for j < l && !IsSpace(buf[j]) {
		j++
	}
	return i, j
}

// NthField return the start and the end index of the nth field in buf.
// n is zero based.
func NthField(buf []byte, n int) (start, end int) {
	i := 0
	offset := 0
	for {
		start, end = NextField(buf)
		if i >= n {
			break
		}
		offset += end + 1
		buf = buf[end+1:]
		i++
	}
	return start + offset, end + offset
}
