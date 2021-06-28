package main

import (
	"fmt"
	"strings"
)

func main() {
	reverseString("hello  abc")
	reverseString("123 4")

	fmt.Println(reverseXin("hello123"))
	fmt.Println(reverseXin("toi la vuong tran"))

	fmt.Println(reverseNew("hello123"))
	fmt.Println(reverseNew("toi la vuong tran"))

}

func reverseString(data string) string {
	slString := strings.Split(data, "")
	fmt.Println(slString)
	reverseSlString := revereSlice(slString)
	fmt.Println(reverseSlString)
	return ""
}

func revereSlice(slString []string) string {
	size := len(slString)
	mod := size / 2
	for i := 0; i < mod; i++ {
		swapStr := slString[i]
		slString[i] = slString[size-1-i]
		slString[size-1-i] = swapStr
	}
	return strings.Join(slString, "")
}

func reverseNew(data string) string {
	split := strings.Split(data, "")
	for i := 0; i < len(split)/2; i++ {
		split[i], split[len(split)-1-i] = split[len(split)-1-i], split[i]
	}

	return strings.Join(split, "")
}

func reverseXin(slString string) string {
	split := []rune(slString)

	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}
	return string(split)
}
