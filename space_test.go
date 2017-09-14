package ascii

import "testing"

func TestIsSpace(t *testing.T) {
	testCases := []struct {
		b    byte
		want bool
	}{
		{b: ' ', want: true},
		{b: '\r', want: true},
		{b: '\n', want: true},
		{b: '\t', want: true},
		{b: '_', want: false},
		{b: '0', want: false},
	}
	for _, c := range testCases {
		got := IsSpace(c.b)
		if got != c.want {
			t.Errorf("failed for %q, got %v, want %v", c.b, got, c.want)
		}
	}
}

func BenchmarkIsSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for c := 0; c <= 255; c++ {
			IsSpace(byte(c))
		}
	}
}
