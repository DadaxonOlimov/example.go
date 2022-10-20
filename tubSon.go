package main

import "fmt"

func main() {
	fmt.Println(tubSon())
}

func tubSon() string {
	var a int
	var count int 
	fmt.Scan(&a)
	if a ==1 {
		return "murakkab"
	}
	for i := 1; i < a; i++ {
		if a % i == 0{
		count += 1
		}
		if  count >= 2 {

		return "murakkab"
	}
	
	}
	return "tub"
}
