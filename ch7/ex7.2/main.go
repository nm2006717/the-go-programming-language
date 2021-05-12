package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

//	练习 7.2：
//	写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
//	返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。

type ByteCounter struct {
	w       io.Writer
	written int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{w, 0}
	return &c, &c.written
}

func main() {
	w, c := CountingWriter(ioutil.Discard)
	fmt.Fprintf(w, "Hello,World!\n")
	fmt.Println(*c)
}
