package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(floor(a, b))

}

func floor(a, b int) int {
	return int(math.Ceil(math.Abs(float64(b-a)) / 10))
}
