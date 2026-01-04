package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(l byte) byte {
	isLetterCapital := l >= 'A' && l <= 'Z'
	isLetterSmall := l >= 'a' && l <= 'z'

	if !isLetterCapital && !isLetterSmall {
		return l // Not a letter
	}

	l += 13 // Move letter 13 positions forward

	isUpdatedLetterOutsideRange := isLetterCapital && l > 'Z' || !isLetterCapital && l > 'z'

	if isUpdatedLetterOutsideRange {
		l -= 26 // Go though an alphabet loop
	}

	return l
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	for i := range b[:n] {
		b[i] = rot13(b[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
