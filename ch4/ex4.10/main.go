package main

import (
	"./github"
	"flag"
	"fmt"
	"log"
	"time"
)

const (
	moreYear    = "moreYear"
	recentYear  = "recentYear"
	recentMonth = "recentMonth"
)

func main() {
	s := flag.String("s", "moreYear", "moreYear:超过一年,recentYear:最近一年,recentMonth:最近一月")
	flag.Parse()

	result, err := github.SearchIssues(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	var beginTime time.Time
	if s != nil {
		switch *s {
		case recentYear:
			beginTime = now.AddDate(-1, 0, 0) //年，月，日   获取一年前的时间
		case recentMonth:
			beginTime = now.AddDate(0, -1, 0) //年，月，日   获取一月前的时间
		default:
			beginTime = now
		}
	}

	// 根据CreateAt筛选result.Items,如果CreateAt<beginTime,移除Item,TotalCount--
	if !beginTime.Equal(now) {
		for i := 0; i < len(result.Items); {
			if result.Items[i].CreatedAt.Before(beginTime) {
				result.Items = append(result.Items[:i], result.Items[i+1:]...)
				result.TotalCount--
				continue
			}
			i++

		}

		fmt.Println(beginTime.Format("2006-01-02 15:04:05"))
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\t%s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}

}
