package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	res := exer1(2000,3500)
	fmt.Println(res)

}

func exer1(low , hight int) string{
	var number []string
	for i := low; i < hight; i++ {
		if i % 7==0 && i % 5 !=0 {
			number = append(number,strconv.Itoa(i))
			
		}
	}
	return strings.Join(number, ",")
}