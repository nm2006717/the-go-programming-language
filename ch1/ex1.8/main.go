package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//	练习 1.8：
//	修改fetch这个范例，如果输入的url参数没有 http:
// 	前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
			os.Exit(1)
		}
	}
}
