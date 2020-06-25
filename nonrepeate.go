package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int{
	lastOccureed := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s){
		if lastI, ok := lastOccureed[ch]; ok && lastI >= start {
			start = lastOccureed[ch] + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccureed[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
			lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
			lengthOfNonRepeatingSubStr("abbb"))


}
