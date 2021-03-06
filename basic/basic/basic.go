package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//作用域-包内部
//var aa = 3
//var ss = "kkk"
//var bb = true
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue()  {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter()  {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler(){
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i * math.Pi)+1)
}

func triangle(){
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts(){
	const (
		filename = "abc.txt"
		a, b = 3, 4
	)
	var c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums()  {
	//const(
	//	cpp = 0
	//	python = 1
	//	golang = 2
	//)
	const(
		cpp = iota
		_
		python
		golang
	)
	const (
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}
