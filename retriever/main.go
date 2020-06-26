package main

import (
	"fmt"
	"golang/retriever/mock"
	"golang/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get("http://www.baidu.com")
}

func inspect(r Retriever){
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)

	}
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "this is a fake test"}

	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	//fmt.Println(download(r))

	//fmt.Println(download(mock.Retriever{Content: "this is a fake test too"}))

	inspect(r)

	// Type assertion
	realRetriver := r.(*real.Retriever)
	fmt.Println(realRetriver.TimeOut)
	if mockRetriver, ok := r.(mock.Retriever); ok{
		fmt.Println(mockRetriver.Content)
	}else {
		fmt.Println("not a mook retriver")
	}

}
