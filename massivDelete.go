package main

import "fmt"

func main()  {
	fmt.Println(delete())
	
}


func delete() int{
	var a int
	fmt.Scan(&a)
	var b []int
	var add int
	for i := 0; i < a; i++ {
		fmt.Scan(&add)
		b = append(b, add)
	}
		
	 maxValue := 0
	index := make(map[int]int)
	for _,v := range b {
		index[v]++
		if index[v] > maxValue{
			maxValue = index[v]
		}

	}
	return a - maxValue
	}
 

