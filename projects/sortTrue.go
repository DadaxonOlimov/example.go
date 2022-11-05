package main

import (
	"fmt"
)

func main()  {
	var count int
	fmt.Scan(&count)
	var numbers []int
	var number int
	for i := 0; i < count; i++ {
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}
	
	fmt.Println(sort(count, numbers))
}

func sort(count int, numbers []int) string {

	for i := 1; i < count; i++ {
		if numbers[i] < numbers[i-1] {
			return "YES"
		}
	}
	return "NO"
}	

