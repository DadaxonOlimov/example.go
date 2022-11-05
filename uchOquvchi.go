package main

import "fmt"

func main()  {
	var number int
	fmt.Scan(&number)
	if 6 < number && number < 7 {
		fmt.Println(9)
	}else{
		fmt.Println(0)
	}
}	