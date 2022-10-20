package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	b := 0
	for i := 0; i <= a; i++ {
		b += i
	}
	fmt.Println(b)

}
