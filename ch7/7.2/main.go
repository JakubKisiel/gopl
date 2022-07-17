package main

import (
	"bufio"
	"fmt"
	"io"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	bufio.ScanWords(p, true)
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type CounterWriter struct {
	writer io.Writer
	count  int64
}

func (w *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{
		writer: w,
		count:  0,
	}
	return cw, &cw.count
}
func main() {
	//!+main
	var c ByteCounter
	cw, counter := CountingWriter(&c)
	cw.Write([]byte("hello"))
	fmt.Println(*counter) // "5", = len("hello")

	*counter = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(cw, "hello, %s", name)
	fmt.Println(*counter) // "12", = len("hello, Dolly")
	//!-main
}
