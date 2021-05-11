package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
//	首先算出x位于哪一个world,然后算出x位于这个word那一位
//	然后在该位置找，看该位置是否为1，若为1则存在，不为1则不存在。
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)

	// word < len(s.words):words中是否存在下标为word。
	// s.words[word]&(1<<bit)，判断该值是否为1
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// 	Add adds the non-negative value x to the set.
// 	首先根据x/64找出这个值位于哪个words
//	然后根据1<<x%64计算出这个值位于words中的哪一位
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 	UnionWith sets s to the union of s and t.
//	求s、t的并集，将结果放在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
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
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))

//	在第一个Println中，我们打印一个*IntSet的指针，这个类型的指针确实有自定义的String方法。第二Println，我们直接调用了x变量的String()方法；
//	这种情况下编译器会隐式地在x前插入&操作符，这样相当于我们还是调用的IntSet指针的String方法。在第三个Println中，因为IntSet类型没有String方法，
//	所以Println方法会直接以原始的方式理解并打印。所以在这种情况下&符号是不能忘的。在我们这种场景下，你把String方法绑定到IntSet对象上，
//	而不是IntSet指针上可能会更合适一些，不过这也需要具体问题具体分析。
	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)
}
