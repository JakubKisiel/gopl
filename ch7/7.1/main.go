package main

import (
	"bufio"
	"fmt"
	"strings"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	bufio.ScanWords(p, true)
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func countOnSplit(p []byte, f bufio.SplitFunc) (int, error) {
	c := 0
	s := string(p)
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(f)
	for sc.Scan() {
		c++
	}
	if err := sc.Err(); err != nil {
		return c, err
	}
	return c, nil
}

//!-bytecounter
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	count, err := countOnSplit(p, bufio.ScanWords)
	*c = WordCounter(count)
	return count, err
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	count, err := countOnSplit(p, bufio.ScanLines)
	*c = LineCounter(count)
	return count, err
}

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
	var wc WordCounter
	wc.Write([]byte("hello"))
	fmt.Println(wc) // "1", = word of "hello"

	wc = 0 // reset the counter
	wtext := "Dolly"
	fmt.Fprintf(&wc, "hello, %s", wtext)
	fmt.Println(wc) // "2", = words of "hello, Dolly"

	var lc LineCounter
	lc.Write([]byte("hello"))
	fmt.Println(lc) // "1", = lines of "hello"

	lc = 0 // reset the counter
	ltext := "Dolly"
	fmt.Fprintf(&lc, "hello,\n %s", ltext)
	fmt.Println(lc) // "2", = lines of "hello,\n Dolly"
}
