package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while fetching url[%s] err:  %v", os.Args[1], err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	fmt.Fprintf(os.Stderr, "error while parsing %v", err)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while parsing %v", err)
		os.Exit(1)
	}
	elmMap := make(map[string]int)
	visit(doc, elmMap)
	for k, v := range elmMap {
		fmt.Printf("%8s -> %d\n", k, v)
	}
}

func visit(n *html.Node, elmMap map[string]int) {
	if n.Type == html.ElementNode {
		elmMap[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, elmMap)
	}
}
