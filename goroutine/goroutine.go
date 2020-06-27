package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i <10; i++{
		go func(i int) {
			for {
				//fmt.Printf("Hello from " + "goroutine %d\n", i)
				a[i]++
				runtime.Gosched()  //手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
