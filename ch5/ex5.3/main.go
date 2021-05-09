package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
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

	texts := visit(doc)
	fmt.Println(texts)

}

func visit(n *html.Node) (texts []string) {
	if n == nil {
		return nil
	}
	if n.Data != "script" && n.Data != "style" {
		texts = append(texts, visit(n.FirstChild)...)
	}
	texts = append(texts, visit(n.NextSibling)...)
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	return
}
