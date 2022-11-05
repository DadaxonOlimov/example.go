package main

import (
	"fmt"
	"strings"
)

func main()  {
	var a string
	fmt.Scan(&a)
	Split := strings.Split(a ,"")
	fmt.Println(Split[0])
}
