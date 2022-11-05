package main

import (
	"fmt"
)


func main()  {
	var a,b int
	fmt.Scan(&a,&b)
	fmt.Println(stones(a,b))
}

func stones(SumalakTosh int, person int) int {
	var add []int
	var count int
	for i := 0; i < person; i++ {
		fmt.Scan(&count)
		add = append(add, count)
	}
	
	for i := 0; i < person; i++ {
		for j := 0; j < person; j++ {
			if add[i] > add[j] {
				add[i], add[j] = add[j], add[i]
			}
		}
	}
	var sum int
	for i, v := range add {
		sum += v
		if SumalakTosh <= sum {
			return i + 1
		}
	}

	return -1
}