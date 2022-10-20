package main

import(
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scan(&a)
	fmt.Println(math.Abs(a*(a-3)/2))
}