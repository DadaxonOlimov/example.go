package main

import (
	"fmt"

)

func main()  {
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	parta := a + b + c
	if parta % 2 == 0 {
		fmt.Println( parta / 2)
	}else  {
		fmt.Println((parta +1) /2)
	}
}