package main

import (
	"bytes"
	"fmt"
)

//	练习6.1:
//	为bit数组实现下面这些方法

type IntSet struct {
	words []uint64
}

// Len return the number of elements
func (s *IntSet) Len() int {
	n := len(s.words)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				cnt++
			}
		}
	}
	return cnt
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, x%64
	return word <= len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, x%64
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 	Remove remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if word == len(s.words)-1 {
		if s.words[word]&(1<<bit) != 0 {
			s.words[word] -= 1 << bit
			if s.words[word] == 0 {
				if len(s.words) == 1 {
					s.words = []uint64{}
				} else {
					s.words = append(s.words[:word])
				}
			}
		}
	}
	if word < len(s.words)-1 {
		if s.words[word]&(1<<bit) != 0 {
			s.words[word] -= 1 << bit
		}
	}
}

//	Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

//	Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var s2 IntSet
	s2 = *s
	return &s2
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

func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	s1 := &IntSet{words: []uint64{}}
	s1.Add(22)
	s1.Add(23)
	s2 := s1.Copy()
	fmt.Println(s2)

	s3 := &IntSet{}
	s3.Add(1)
	s3.Add(2)
	s3.Add(5)
	s3.Clear()
	s3.Remove(1)

	fmt.Println(s3)
	fmt.Println(s3.Len())
}
