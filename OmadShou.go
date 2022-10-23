package main

import "fmt"

func main()  {
	var bottle, liter,try int
	fmt.Scan(&bottle,&liter,&try)
	result := bottle *(liter -1)
	if try >= result + 1{
		fmt.Println(1)
	}else {
		fmt.Println(-1)
	}
	
}