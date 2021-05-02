package main

// 	练习 4.7：
//	修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
