package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//	中序打印出树
func (t *tree) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	var helper func(node *tree)
	helper = func(node *tree) {
		if node == nil {
			return
		}
		helper(node.left)
		//	控制打印空格
		// 	第一个元素前不用打印空格
		if buf.Len() > len("[") {
			buf.WriteByte(' ')
		}
		buf.WriteString(strconv.Itoa(node.value))
		helper(node.right)
	}
	helper(t)
	buf.WriteByte(']')
	return buf.String()
}



func main() {
	t := &tree{}
	t.value = 1
	t.left = &tree{
		value: 10,
		left:  nil,
		right: nil,
	}
	t.right = &tree{
		value: 11,
		left:  nil,
		right: nil,
	}

	fmt.Println(t)

}
