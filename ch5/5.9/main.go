package main

import (
	"fmt"
	"log"
	"strings"
)

const toReplace = "123321$fooooo$foo$foodsadsa"

func main() {
	fmt.Println(toReplace)
	replaced := expand(toReplace, changeAscii)
	fmt.Println(replaced)
}

func changeAscii(s string) string {
	return strings.Map(func(r rune) rune {
		return r + 1
	}, s)
}

func expand(s string, f func(string) string) string {
	if f == nil {
		log.Fatalf("Passed function %T cannot be nil", f)
	}
	replacement := f("foo")
	replacer := strings.ReplaceAll(s, "$foo", replacement)
	return replacer
}
