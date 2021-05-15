package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//	练习 7.9：
//	使用html/template包 (§4.6) 替代printTracks将tracks展示成一个HTML表格。
//	将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

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
	Tracks []*Track
	less   func(x *Track, y *Track) bool
	clicks []string
}

func (c customTrack) Len() int {
	return len(c.Tracks)
}

func (c customTrack) Less(i, j int) bool {
	return c.less(c.Tracks[i], c.Tracks[j])
}

func (c customTrack) Swap(i, j int) {
	c.Tracks[i], c.Tracks[j] = c.Tracks[j], c.Tracks[i]
}

var customT = customTrack{
	Tracks: tracks,
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

var musicList = template.Must(template.New("musicList").Parse(`
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>Artist</th>
  <th>Album</th>
  <th>Year</th>
  <th>Length</th>
</tr>
{{range .Tracks}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {

	http.HandleFunc("/music", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html;charset=utf-8")
		if err := request.ParseForm(); err != nil {
			log.Print(err)
		}
		params := request.Form["q"]
		if len(params) != 0 {
			pop(&customT.clicks)
			push(&(customT.clicks), params[0])
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
		if err := musicList.Execute(writer, customT); err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe("localhost:8000", nil)

}
