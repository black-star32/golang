package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int){
	for {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}

func createWorker(id int)chan<- int{
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo()  {
	//var c chan int // c == nil
	//c := make(chan int)
	//go func() {
	//	for {
	//		n := <- c
	//		fmt.Println(n)
	//	}
	//}()


	//var channels [10]chan int
	//for i := 0; i < 10; i++{
	//	channels[i] = make(chan int)
	//	go worker(i, channels[i])
	//}
	//
	//for i := 0; i < 10; i++{
	//	channels[i] <- 'a' + i
	//}
	//
	//for i := 0; i < 10; i++{
	//	channels[i] <- 'A' + i
	//}

	var channels [10]chan<- int
	for i := 0; i < 10; i++{
		channels[i] = createWorker(i)
		//go worker(i, channels[i])\]
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func worker2(id int, c chan int){
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d received %c\n", id, n)
	//}
	for n := range c{
		fmt.Printf("worker %d received %c\n", id, n)
	}
}

func createWorker2(id int)chan<- int{
	c := make(chan int)
	go worker2(id, c)
	return c
}

func channleClose()  {
	c := make(chan int)
	go worker2(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker2(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channleClose()
}


