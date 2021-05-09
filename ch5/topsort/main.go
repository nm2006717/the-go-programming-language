package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
//var prereqs = map[string][]string{
//	"algorithms": {"data structures"},
//	"calculus":   {"linear algebra"},
//	"compilers": {
//		"data structures",
//		"formal languages",
//		"computer organization",
//	},
//	"data structures":       {"discrete math"},
//	"databases":             {"data structures"},
//	"discrete math":         {"intro to programming"},
//	"formal languages":      {"discrete math"},
//	"networks":              {"operating systems"},
//	"operating systems":     {"data structures", "computer organization"},
//	"programming languages": {"data structures", "computer organization"},
//}

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

// 递归匿名函数写法
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
