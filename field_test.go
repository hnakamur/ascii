package ascii

import (
	"bytes"
	"testing"
)

func TestNextField(t *testing.T) {
	buf := []byte("MemTotal:       16260508 kB\n")
	wants := [][]byte{[]byte("MemTotal:"), []byte("16260508"), []byte("kB")}
	i := 0
	for len(buf) > 0 {
		start, end := NextField(buf)
		got := buf[start:end]
		want := wants[i]
		if !bytes.Equal(got, want) {
			t.Errorf("field %d unmatch, got %q, want %q", i, got, want)
		}
		i++
		buf = buf[end+1:]
	}
}

func TestNthField(t *testing.T) {
	buf := []byte("MemTotal:       16260508 kB\n")
	testCases := []struct {
		n    int
		want []byte
	}{
		{0, []byte("MemTotal:")},
		{1, []byte("16260508")},
		{2, []byte("kB")},
		{3, []byte{}},
	}
	for _, c := range testCases {
		start, end := NthField(buf, c.n)
		got := buf[start:end]
		want := c.want
		if !bytes.Equal(got, want) {
			t.Errorf("field %d unmatch, got %q, want %q", c.n, got, want)
		}
	}
}

func BenchmarkNextField(b *testing.B) {
	buf := []byte("MemTotal:       16260508 kB\n")
	for i := 0; i < b.N; i++ {
		for len(buf) > 0 {
			_, end := NextField(buf)
			buf = buf[end+1:]
		}
	}
}
