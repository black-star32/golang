package main

import (
	"fmt"
	"golang/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpyt())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpyt())

	q.Push("abc")
	fmt.Println(q.Pop())
}
