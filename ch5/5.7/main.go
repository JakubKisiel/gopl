package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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
	depth := 0
	visit(doc, &depth, beforeElement, afterElement)
}

func beforeElement(n *html.Node, depth *int) {
	if n.Type == html.TextNode && n.Data != "" {
		fmt.Printf("%*s%s\n", *depth*2, "", n.Data)
	}
	if n.Type != html.ElementNode {
		return
	}
	var sb strings.Builder
	sb.WriteString(n.Data)

	for _, attr := range n.Attr {
		attrStr := fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
		sb.WriteString(attrStr)
	}
	if n.FirstChild == nil {
		sb.WriteString("/")
	}
	fmt.Printf("%*s<%s>\n", *depth*2, "", sb.String())
	*depth++
}

func afterElement(n *html.Node, depth *int) {
	if n.Type != html.ElementNode {
		return
	}
	*depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Printf("%*s</%s>\n", *depth*2, "", n.Data)
}

func visit(n *html.Node, depth *int, pre, post func(*html.Node, *int)) {
	if pre != nil {
		pre(n, depth)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		visit(c, depth, pre, post)
	}
	if post != nil {
		post(n, depth)
	}
}
