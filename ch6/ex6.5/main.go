package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//	练习 6.5：
//	我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。
//	当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
type IntSet struct {
	words []uint
}

/*
	在strconv包中有个常量
	const intSize  =  32  <<  (  ^uint(0)  >>  63  )
	const IntSize = intSize
	在32位平台这个值为32;在64位平台,这个值为64; 因此可以通过这个来判断平台的位数.
原理:
	在32平台系统:
	1. uint(0)在平台底层是0x00000000
	2. ^uint(0)在平台底层是0xFFFFFFFF
	3. ^uint(0) >> 63 在底层平台是0x00000000,也就是0
	4. 32 << 0 结果是32

	在64平台系统:
	1. uint(0)在平台底层是0x0000000000000000
	2. ^uint(0)在平台底层是0xFFFFFFFFFFFFFFFF
	3. ^uint(0) >> 63 在底层平台是0x0000000000000001,也就是1
	4. 32 << 1 结果是32*2  =  64
*/
const intSize = strconv.IntSize

// Len return the number of elements
func (s *IntSet) Len() int {
	n := len(s.words)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < intSize; j++ {
			if s.words[i]&(1<<j) != 0 {
				cnt++
			}
		}
	}
	return cnt
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/intSize, x%intSize
	return word <= len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/intSize, x%intSize
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 	Remove remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/intSize, x%intSize
	if word == len(s.words)-1 {
		if s.words[word]&(1<<bit) != 0 {
			s.words[word] -= 1 << bit
			if s.words[word] == 0 {
				if len(s.words) == 1 {
					s.words = []uint{}
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
	s.words = []uint{}
}

//	Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var s2 IntSet
	s2 = *s
	return &s2
}

// 	UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//	IntersectWith 计算s、t两个集合的交集，结果放在s上
func (s *IntSet) IntersectWith(t *IntSet) {
	nS := len(s.words)
	nT := len(t.words)
	if nS == 0 || nT == 0 {
		s.words = []uint{}
	}
	if nS > nT {
		s.words = s.words[0:nT]
	} else {
		t.words = t.words[0:nS]
	}
	for i := nT - 1; i >= 0; i-- {
		s.words[i] &= t.words[i]
		if s.words[i] == 0 {
			s.words = s.words[0:i]
		}
	}
}

//	DifferenceWith 计算s,t两个集合的差集（元素出现在s中，且没有出现在t中），结果放在s上。
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := 0; i < len(t.words); i++ {
		s.words[i] &= ^t.words[i]
	}
}

func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//	AddALl 添加一组IntSet
func (s *IntSet) AddAll(nums ...int) {
	if len(nums) != 0 {
		for _, v := range nums {
			s.Add(v)
		}
	}
}

func (s *IntSet) Elems() []uint {
	var elems []uint
	for i := 0; i < len(s.words); i++ {
		for j := 0; j < intSize; j++ {
			if s.words[i]&(1<<j) != 0 {
				elems = append(elems, uint(intSize*i+j))
			}
		}
	}
	return elems
}

func main() {
	s1 := &IntSet{words: []uint{}}
	s1.AddAll(22,23)
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

	s4 := &IntSet{}
	s4.AddAll([]int{1, 2, 3, 4, 5, 6}...)

	fmt.Println(s4)

	s5 := &IntSet{}
	s6 := &IntSet{}
	s5.AddAll(1, 2)
	s6.AddAll(3, 4, 5)

	s5.IntersectWith(s6)
	fmt.Println(s5)

	s7 := &IntSet{}
	s7.AddAll(1, 2, 3, 4, 5, 6, 7, 8)
	s8 := &IntSet{}
	s8.AddAll(4, 5, 6, 7)
	s7.DifferenceWith(s8)
	fmt.Println(s7)

	s9 := &IntSet{}
	s10 := &IntSet{}
	s9.DifferenceWith(s10)
	fmt.Println(s9)

	s11 := &IntSet{}
	s11.AddAll(11, 12, 14)

	elems := s11.Elems()

	fmt.Println(elems)

}
