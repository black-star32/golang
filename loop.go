package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string{
	result := ""
	for ; n > 0; n /=2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return  result
}

func pritFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil{
		println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func forver()  {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		converToBin(5), // 101
		converToBin(13), // 1011 --> 1101
		converToBin(772323),
		converToBin(0),
		)

	pritFile("abc.txt")
	//forver()
}
