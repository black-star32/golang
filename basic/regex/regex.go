package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com
email1 is abc@def.org
email2 is kkk@qq.com
email3 is dddd@qq.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)

	for _, m := range match{
		fmt.Println(m)
	}
}
