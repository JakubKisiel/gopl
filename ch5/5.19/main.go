package main

import "fmt"

func main() {
	fmt.Println(yolo())
}
func yolo() (s string) {
	defer func() {
		recover()
		s = "XD"
	}()
	panic("123")
}
