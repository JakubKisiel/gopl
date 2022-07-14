package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println(max(54324, 213213121, 87698679, 34124312))
	fmt.Println(min(543, 5432, 321, -5423542, -78675, 4))
}

func varInts(f func(int, int) int, ints ...int) int {
	if len(ints) == 0 {
		return 0
	}
	if f == nil {
		log.Fatalf("Compartor %T cannot be nil ", f)
	}
	max := ints[0]
	for _, i := range ints {
		max = f(max, i)
	}
	return max
}

func max(ints ...int) int {
	return varInts(func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}, ints...)
}

func min(ints ...int) int {
	return varInts(func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}, ints...)
}

func varJoin(j string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(strs[0])
	for _, s := range strs[1:] {
		sb.WriteString(s)
		sb.WriteString(j)
	}
	return sb.String()
}
