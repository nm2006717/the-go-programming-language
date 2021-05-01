package main

import (
	"./github"
	"html/template"
	"log"
	"net/http"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {

	http.HandleFunc("/issue", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := request.ParseForm(); err != nil {
			log.Print(err)
		}
		params := request.Form["q"]
		if len(params) == 0 {
			writer.Write([]byte("<b>url查询参数出错</b>"))
		}else{
			result, err := github.SearchIssues(params)
			if err != nil {
				log.Fatal(err)
			}
			if err := issueList.Execute(writer, result); err != nil {
				log.Fatal(err)
			}
		}

	})

	http.ListenAndServe("localhost:8000", nil)

}
