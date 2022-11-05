package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
	fmt.Scan(&str1)
	res1 := strings.Split(str1, "")
	fmt.Println(res1)
	if len(res1) <= 2 {
		fmt.Println(str1)
	} else if res1[0] > res1[len(res1)-1] {
		fmt.Println(res1)
	} else {
		fmt.Println("sassiq")
	}

}

func Buratino(str1 string) {
	fmt.Scan(&str1)
	res1 := strings.Split(str1, "")
	fmt.Println(res1)
	if len(res1) <= 2 {
		fmt.Println(str1)
	} else if res1[0] > res1[len(res1)-1] {
		fmt.Println(res1[len(res1)-1])
	}

}
