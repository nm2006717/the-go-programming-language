package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//	练习 7.12： 修改/list的handler让它把输出打印成一个HTML的表格而不是文本。
//	html/template包(§4.6)可能会对你有帮助。
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/insert", db.insert)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

var goodsList = template.Must(template.New("goodsList").Parse(`
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $key, $value := .}}
<tr>
  <td>{{$key}}</td>
  <td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

type dollars float32

var mu sync.Mutex

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	//for item, price := range db {
	//	fmt.Fprintf(w, "%s: %s\n", item, price)
	//}
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if err := goodsList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	defer mu.Unlock()
	_, ok := db[item]
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "delete item:%q,price:%q\n", item, db[item])
}
func (db database) insert(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		http.Error(w, fmt.Sprintf("params incorrect! "), http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	_, ok := db[item]
	if ok {
		http.Error(w, fmt.Sprintf("has item: %q already", item), http.StatusNotFound)
		return
	}
	pri, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("server internal error"), http.StatusInternalServerError)
		return
	}
	db[item] = dollars(pri)
	fmt.Fprintf(w, "insert success,item:%q price:%q\n", item, db[item])
}
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	_, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "no such item: %q ", item)
		return
	}
	pri, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("server internal error"), http.StatusInternalServerError)
		return
	}
	db[item] = dollars(pri)

	fmt.Fprintf(w, "update success,item:%q price:%q\n", item, db[item])
}
