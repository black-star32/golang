package main

import (
	"bufio"
	"fmt"
	"golang/functional/fib"
	"io"
	"strings"
)

func fibonacci() intGen{
	a, b := 0, 1
	println(a, b)
	return func() int {
		a, b = b, a + b
		return a
	}
}

type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000{
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	//f := fibonacci()

	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 2
	//fmt.Println(f()) // 3
	//fmt.Println(f()) // 5
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	//printFileContents(f)

	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
