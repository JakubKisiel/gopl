package main

import (
	"fmt"
	"net/http"
	"os"

	"gopl/ch5/5.7/traverse"

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
	//base padding is 2 if u want no padding provide -2 value
	depth := 0
	traverse.Visit(doc, &depth, traverse.BeforeElement, traverse.AfterElement)
}
