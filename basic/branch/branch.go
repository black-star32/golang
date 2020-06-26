package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func grade(score int) string{
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}


func main() {
	const filename = "basic/branch/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(75),
		grade(83),
		//grade(101),
		)
	s := `sefasf
sdfsdf
"dasf"`

	printFileContents(strings.NewReader(s))
}
