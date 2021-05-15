package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//	练习 7.8：
//	很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次点击过列头的列，
//	第二个排序键是第二最近点击过列头的列，等等。定义一个sort.Interface的实现用在这样的表格中。
//	比较这个实现方式和重复使用sort.Stable来排序的方式。

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() //	calculate column widths and print table
}

type customTrack struct {
	tracks []*Track
	less   func(x *Track, y *Track) bool
	clicks []string
}

func (c customTrack) Len() int {
	return len(c.tracks)
}

func (c customTrack) Less(i, j int) bool {
	return c.less(c.tracks[i], c.tracks[j])
}

func (c customTrack) Swap(i, j int) {
	c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i]
}

var customT = customTrack{
	tracks: tracks,
	less:   nil,
	clicks: []string{"title", "artist", "album", "year", "length"},
}

func push(clicks *[]string, click string) {
	result := make([]string, len(*clicks)+1)
	result[0] = click
	copy(result[1:], *clicks)
	copy(*clicks, result)
}

func pop(clicks *[]string) {
	*clicks = append((*clicks)[:len(*(clicks))-1])
}

func main() {

	click := flag.String("click", "", "点击排序")
	flag.Parse()
	if click != nil && *click != "" {
		pop(&customT.clicks)
		push(&(customT.clicks), *click)
	}
	customT.less = func(x *Track, y *Track) bool {
		for i := 0; i < len(customT.clicks); i++ {
			switch customT.clicks[i] {
			case "title":
				if x.Title != y.Title {
					return x.Title < y.Title
				}
			case "artist":
				if x.Artist != y.Artist {
					return x.Artist < y.Artist
				}
			case "album":
				if x.Album != y.Album {
					return x.Album < y.Album
				}
			case "year":
				if x.Year != y.Year {
					return x.Year < y.Year
				}
			case "length":
				if x.Length != y.Length {
					return x.Length < y.Length
				}
			}
		}
		return false
	}

	sort.Sort(customT)

	printTracks(customT.tracks)

}
