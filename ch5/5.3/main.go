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
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while parsing %v", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("Printing contents of %v \nContent: %s\n", n.Parent.Data, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		visit(c)
	}
}
