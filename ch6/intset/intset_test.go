package intset

import (
	"fmt"
	"strconv"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	type Test struct {
		arr    []int
		result int
	}
	testCases := []Test{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 43243, 65436543, 43211}, 4},
	}

	for i, testCase := range testCases {
		testFunc := func(t *testing.T) {
			var x IntSet
			for _, val := range testCase.arr {
				x.Add(val)
			}
			if x.Len() != testCase.result {
				t.Errorf("Expected: %d\nto be equal to: %d\nIntSet: %s",
					x.Len(), testCase.result, x.String())
			}
		}
		t.Run(strconv.Itoa(i), testFunc)
	}
}
func TestRemove(t *testing.T) {
	type Test struct {
		arr    []int
		remove int
	}
	testCases := []Test{
		{[]int{}, 3},
		{[]int{1}, 1},
		{[]int{1, 3, 65436543, 43211}, 3},
	}

	for i, testCase := range testCases {
		testFunc := func(t *testing.T) {
			var x IntSet
			for _, val := range testCase.arr {
				x.Add(val)
			}
			x.Remove(testCase.remove)
			if x.Has(testCase.remove) {
				t.Errorf("Expected: %d \nto be absent in:%s",
					testCase.remove, x.String())
			}
		}
		t.Run(strconv.Itoa(i), testFunc)
	}
}

func TestCopy(t *testing.T) {
	type Test struct {
		arr []int
	}
	testCases := []Test{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 3, 65436543, 43211}},
	}

	for i, testCase := range testCases {
		testFunc := func(t *testing.T) {
			var x IntSet
			for _, val := range testCase.arr {
				x.Add(val)
			}
			n := x.Copy()
			for _, result := range testCase.arr {
				if !n.Has(result) {
					t.Errorf("Expected: %s \nto be equal to: %s",
						n.String(), x.String())
				}
			}
		}
		t.Run(strconv.Itoa(i), testFunc)
	}
}
