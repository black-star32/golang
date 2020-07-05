package main

import (
	"golang/crawler/frontend/controller"
	"net/http"
)

//http://localhost:8888/search?q=%E7%94%B7%E5%A3%AB&from=1
func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(""))
	err := http.ListenAndServe(":8888", nil)
	if err != nil{
		panic(err)
	}
}
