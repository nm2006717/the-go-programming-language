package popcount

//	练习 2.3： 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。
//	（11.4节将展示如何系统地比较两个不同实现的性能。）

// 	解答:利用x&(x-1)会将x最低位的1替换为0的性质。
// 	比如: x:			1 0 0 1
//		x-1: 		1 0 0 0
//      x&(x-1) 	1 0 0 0

//      x= x&(x-1)	1 0 0 0
//		x-1			0 1 1 1
//		x&(x-1)		0 0 0 0

//		return 2
func PopCount(x uint64) int {
	ret := 0
	for x != 0 {
		x = x & (x-1)
		ret++
	}
	return ret
}
