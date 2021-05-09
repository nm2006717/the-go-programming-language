package main

import "fmt"

//	练习5.10：
//	重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。

var prereqs = map[string][]string{
	"算法":  {"数据结构"},
	"微积分": {"线性代数"},
	"编译原理": {
		"数据结构",
		"形式语言 ",
		"计算机体系结构",
	},
	"数据结构":   {"离散数学"},
	"数据库":    {"数据结构"},
	"离散数学":   {"编程导论"},
	"形式语言":   {"离散数学"},
	"网络":     {"操作系统"},
	"操作系统":   {"数据结构", "计算机体系结构"},
	"程序设计语言": {"数据结构", "计算机体系结构"},
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[len(order)] = item
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}

func main() {
	maps := topoSort(prereqs)
	for i := 0; i < len(maps); i++ {
		fmt.Printf("%d:%s\n", i, maps[i])
	}

}
