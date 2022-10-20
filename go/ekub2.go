package main

import (
	"fmt"
	
)

func main() {
	var a[]int
	fmt.Scan(&a)
	as :=ranges(12)
	fmt.Println(as)
}

func ranges(s int)map[int]int{
	a :=make(map[int]int)

	fmt.Scan(a)
	for i, ok := range a {
	if a[i] == 0 {
		continue
	}
	if a[i] < a[i-1]{

		fmt.Println("not")
	}
		//return false
}
return a
}