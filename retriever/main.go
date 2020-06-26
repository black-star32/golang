package main

import (
	"fmt"
	"golang/retriever/mock"
	"golang/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "this is a fake test"}
	r = real.Retriever{}
	fmt.Println(download(r))

	//fmt.Println(download(mock.Retriever{Content: "this is a fake test too"}))
}
