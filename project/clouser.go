package main

import "fmt"

func dosomthing(colback func(int ,int) int, s string) int  {
	result := colback(5,11)
	fmt.Println(s)
	return result
	
}

func main()  {
	sumColback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := dosomthing(sumColback, "plus")
	fmt.Println(result)
		
	
}