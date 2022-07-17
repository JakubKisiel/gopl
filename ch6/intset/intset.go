package intset

import (
	"bytes"
	"fmt"
)

//size for intset to be more efficient on 32-bit platforms
const uint_size = 32 << (^uint(0) >> 63)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uint_size, uint(x%uint_size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uint_size, uint(x%uint_size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uint_size; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uint_size*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	counter := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uint_size; j++ {
			if word&(1<<uint(j)) == 0 {
				continue
			}
			counter++
		}
	}
	return counter
}

func (s *IntSet) Remove(x int) {
	word, bit := x/uint_size, uint(x%uint_size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] ^= 1 << bit & s.words[word]
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var x IntSet
	x.words = append(x.words, s.words...)
	return &x
}

func (s *IntSet) AddAll(vars ...int) {
	for _, val := range vars {
		s.Add(val)
	}
}
