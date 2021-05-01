package main

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

}

// appendValues appends the elements of t to values in order
// and return the resulting slice
// 中序遍历
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = appendValues(values, t.right)
		values = appendValues(values, t.right)
	}
	return values
}

//将数组的值，插入到树种
func add(t *tree, value int) *tree {
	if t == nil {
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
