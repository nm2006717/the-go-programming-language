package popcount

//	练习 2.3： 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。
//	（11.4节将展示如何系统地比较两个不同实现的性能。）
func PopCount(x uint64) int {
	ret := 0
	for x != 0 {
		if x&1 == 1 {
			ret++
		}
		x = x >> 1
	}
	return ret

}
