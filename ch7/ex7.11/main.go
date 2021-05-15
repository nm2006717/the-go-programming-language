package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//	练习 7.11： 增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。
//	例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。
//	（注意：这个修改会引入变量同时更新的问题）
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/insert", db.insert)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

type dollars float32

var mu sync.Mutex

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
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
