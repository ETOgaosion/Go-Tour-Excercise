package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	for i := 0; i < n; i++ {
		if (b[i] < 'a' || b[i] > 'z') && (b[i] < 'A' || b[i] > 'Z') {
			continue
		}
		trans := 'n' - 'a'
		if (b[i] <= 'z' && b[i] > 'm') || (b[i] <= 'Z' && b[i] > 'M') {
			trans = -trans
		}
		b[i] += byte(trans)
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
