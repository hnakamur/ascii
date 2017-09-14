package ascii

import (
	"bytes"
	"testing"
)

func TestGetLine(t *testing.T) {
	buf := []byte("line1\nline2\nline3")
	wants := [][]byte{[]byte("line1\n"), []byte("line2\n"), []byte("line3")}
	i := 0
	p := buf
	for len(p) > 0 {
		got := GetLine(p)
		want := wants[i]
		if !bytes.Equal(got, want) {
			t.Errorf("line %d unmatch, got %q, want %q", i, got, want)
		}
		i++
		p = p[len(got):]
	}
}

func BenchmarkGetLine(b *testing.B) {
	buf := []byte("line1\nline2\nline3")
	for i := 0; i < b.N; i++ {
		p := buf
		for len(p) > 0 {
			got := GetLine(p)
			p = p[len(got):]
		}
	}
}
