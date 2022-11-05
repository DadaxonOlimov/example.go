package main

import (
	"fmt"
)

func main()  {
	var Son int                                                                                      
	var allTime = 50
	fmt.Scan(&Son)
	result := Son % allTime
	if result <= 25 {
		fmt.Println("0__")
	} else if 25 < result && result <= 29 {
		fmt.Println("00_")
	} else if 29 < result && result <= 35 {
		fmt.Println("_0_")
	} else {
		fmt.Println("__0")
	}
}



