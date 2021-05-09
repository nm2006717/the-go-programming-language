package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

const (
	image  = "img"
	script = "script"
	style  = "style"
	other  = "other"
)

//	练习 5.3：
//	编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
func main() {
	resp, err := http.Get("https://geektutu.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "get geektutu: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find text element: %v\n", err)
		os.Exit(1)
	}

	nodesMap := make(map[string][]string)
	nodesMap[image] = make([]string, 0)
	nodesMap[style] = make([]string, 0)
	nodesMap[script] = make([]string, 0)
	nodesMap[other] = make([]string, 0)

	visit(doc, nodesMap)
	fmt.Println(nodesMap)

}

func visit(n *html.Node, nodesMap map[string][]string) () {
	if n == nil {
		return
	}

	visit(n.FirstChild, nodesMap)
	visit(n.NextSibling, nodesMap)

	if n.Type == html.ElementNode {
		switch n.Data {
		case image:
			nodesMap[image] = append(nodesMap[image], n.Data)
		case style:
			nodesMap[style] = append(nodesMap[style], n.Data)
		case script:
			nodesMap[script] = append(nodesMap[script], n.Data)
		default:
			nodesMap[other] = append(nodesMap[other], n.Data)
		}
	}

	return
}
