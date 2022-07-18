package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	p := make([]byte, 5)
	reader := strings.NewReader("The Go Programming language")
	lr := LimitReader(reader, 12)
	for {
		b, err := lr.Read(p)
		if b > 0 {
			fmt.Println(string(p[:b]))
		}
		if err != nil {
			break
		}
	}
}

type limitReader struct {
	r io.Reader
	n int
}

func (r *limitReader) Read(p []byte) (int, error) {
	if len(p) >= r.n {
		p = p[:r.n]
	}
	n, err := r.r.Read(p)
	r.n -= n
	if err != nil {
		return n, err
	}
	if r.n <= 0 {
		return n, io.EOF
	}
	return n, nil
}
func LimitReader(r io.Reader, n int) io.Reader {
	return &limitReader{r, n}
}
