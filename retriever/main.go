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

const url = "http://www.baidu.com"
func download(r Retriever) string{
	return r.Get(url)
}

func inspect(r Retriever){
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Println("> Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
	fmt.Println()
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster){
	poster.Post(url,
		map[string]string{
		"name":"ccmouse",
		"course":"golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string{
	s.Post(url, map[string]string{
		"content": "annoter fake com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{Content: "this is a fake test"}

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
	if mockRetriver, ok := r.(*mock.Retriever); ok{
		fmt.Println(mockRetriver.Content)
	}else {
		fmt.Println("not a mook retriver")
	}

	fmt.Println("try a session")
	retriever := mock.Retriever{Content: "this is a fake test"}
	fmt.Println(session(&retriever))
}
