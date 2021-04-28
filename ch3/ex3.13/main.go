package main

import "fmt"

//	练习 3.13：
//	编写KB、MB的常量声明，然后扩展到YB。

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	KB = 1000
	MB = MiB - MiB%(KB*KB)
	GB = GiB - GiB%(MB*KB)
	TB = TiB - TiB%(GB*KB)
	PB = PiB - PiB%(TB*KB)
	EB = EiB - EiB%(PB*KB)
	ZB = ZiB - ZiB%(EB*KB)
	YB = YiB - YiB%(ZB*KB)
)

func main() {
	fmt.Println(MiB)
}
