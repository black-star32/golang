package main

import (
	"fmt"
	"sync"
)

//func doWork(id int, c chan int, done chan bool){
//	for n := range c {
//		fmt.Printf("worker %d received %c\n", id, n)
//		go func() {done <- true}()
//	}
//}
//func doWork(id int, c chan int, done chan bool){
//	for n := range c {
//		fmt.Printf("worker %d received %c\n", id, n)
//		done <- true
//	}
//}
func doWork(id int, c chan int, wg *sync.WaitGroup){
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		wg.Done()
	}
}

//type worker struct {
//	in chan int
//	done chan bool
//}
type worker struct {
	in chan int
	wg *sync.WaitGroup
}

//func createWorker(id int) worker{
//	w := worker{
//		in: make(chan int),
//		done:make(chan bool),
//	}
//	go doWork(id, w.in, w.done)
//	return w
//}

func createWorker(id int, wg *sync.WaitGroup) worker{
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, w.wg)
	return w
}

func chanDemo()  {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++{
		workers[i] = createWorker(i, &wg)
	}


	wg.Add(20)
	for i, worker := range workers{
		worker.in <- 'a' + i
	}

	//for _, worker := range workers {
	//	<-worker.done
	//}

	for i, worker := range workers{
		worker.in <- 'A' + i
	}

	wg.Wait()

	//time.Sleep(time.Millisecond)

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
	//for _, worker := range workers {
	//	<-worker.done
	//}
}


func main() {
	chanDemo()
}


