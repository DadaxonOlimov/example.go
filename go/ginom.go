package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, e, f, v, k float64
	fmt.Scan(&a, &b, &c, &d, &e, &f, &v, &k)
	ginom := a + b + c + d + e + f + v
	fmt.Println(math.Abs(k - ginom))

}
