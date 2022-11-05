package main

import (
	"fmt"
	
)

func main()  {
	var a int
	var b int
	fmt.Scan(&a)
	var c []int
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		c = append(c, b)
	}
	for i := 0; i < a; i++ {
		for j := 0;  j < a; j++ {
			if c[i] > c[j] {
				c[i],c[j] = c[j],c[i]
			}
		}
	}
	fmt.Println(c[1])
}