package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd *rot13Reader) Read(b []byte) (int, error) {
	length, e := rd.r.Read(b)
	n := 0
	for ; n < length; n++ {
		switch {
		case b[n] >= 'A' && b[n] <= 'Z':
			b[n] = (b[n]-'A'+13)%26 + 'A'
		case b[n] >= 'a' && b[n] <= 'z':
			b[n] = (b[n]-'a'+13)%26 + 'a'
		}
	}
	return length, e
}

func main() {
	s := strings.NewReader("Rira_lbh_Oehghf?")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
