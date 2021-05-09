//	练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，
// 	中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，
// 	根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

var (
	testI   int
	depth   int
	problem string = "tto"
)

func main() {
	s := "<p id='tto'>ddd<a>aaaaa</a></p><p>cccc</p>"
	read := strings.NewReader(s)
	doc, err := html.Parse(read)
	if err != nil {
		fmt.Println(err)
	}
	visit(doc, start, end)
}
func visit(n *html.Node, start, end func(n *html.Node) bool) {
	if er := start(n); !er {
		return
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		visit(i, start, end)
	}
	if er := end(n); !er {
		return
	}
}

func start(n *html.Node) bool {
	switch n.Type {
	case html.ElementNode:
		var str string
		for _, v := range n.Attr {
			str = fmt.Sprintf("<%s %s='%s'></%s>", n.Data, v.Key, v.Val, n.Data)
		}
		for _, v := range n.Attr {
			if v.Key == "id" {
				if v.Val == problem {
					fmt.Println(str)
					os.Exit(0)
				}

			}
		}

		fmt.Printf("%*s<%s>%d\n", depth*2, "", n.Data, depth)
		depth++
	case html.TextNode:
		fmt.Printf("%*s%s%d\n", depth*2, "", n.Data, depth)
	case html.CommentNode:
		fmt.Println(n.Data)
	case html.DoctypeNode:
		fmt.Println(n.Data)
	case html.ErrorNode:
		fmt.Println(n.Data)
	case html.DocumentNode:
		fmt.Println(n.Data)
	}
	return true
}
func end(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>%d\n", depth*2, "", n.Data, depth)
	}
	return true
}