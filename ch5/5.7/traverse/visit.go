package traverse

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func BeforeElement(n *html.Node, depth *int) {
	if n.Type == html.TextNode && n.Data != "" {
		fmt.Printf("%*s%s\n", *depth*2, "", n.Data)
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<--%s-->\n", *depth*2, "", n.Data)
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

func AfterElement(n *html.Node, depth *int) {
	if n.Type != html.ElementNode {
		return
	}
	*depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Printf("%*s</%s>\n", *depth*2, "", n.Data)
}

func Visit(n *html.Node, depth *int, pre, post func(*html.Node, *int)) {
	if pre != nil {
		pre(n, depth)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		Visit(c, depth, pre, post)
	}
	if post != nil {
		post(n, depth)
	}
}
