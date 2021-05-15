package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//	练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
//	并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数：

type LimitedReader struct {
	R io.Reader // underlying reader
	N int64     // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {

	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{R: r, N: n}
}

func main() {
	lr := LimitReader(strings.NewReader("abc"), 1)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%v", err)
	}
	fmt.Printf("%s\n", b)
}