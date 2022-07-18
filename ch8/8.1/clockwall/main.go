package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type Clock struct {
	name string
	addr string
}

func (c *Clock) startPrinting() {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		time := sc.Text()
		fmt.Printf("%s: %s\n", c.name, time)
	}
}

func main() {
	clocks := make([]Clock, 0)
	var wg sync.WaitGroup
	for _, a := range os.Args[1:] {
		kv := strings.Split(a, "=")
		clocks = append(clocks, Clock{kv[0], kv[1]})
	}
	for _, c := range clocks {
		wg.Add(1)
		go func(c Clock) {
			defer wg.Done()
			c.startPrinting()
		}(c)
	}
	wg.Wait()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
