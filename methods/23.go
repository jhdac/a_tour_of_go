package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b uint8) uint8 {
	var final uint8 = b
	switch {
	case b >= 'A' && b <= 'Z':
		final = b - 'A' + 13
		final %= 26
		final += 'A'
	case b >= 'a' && b <= 'z':
		final = b - 'a' + 13
		final %= 26
		final += 'a'
	}
	return final
}

func (r rot13Reader) Read2(b []byte) (int, error) {
	n, err := r.r.Read(b)
	return n, err
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}

	i := 0
	for ; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return i, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}