package main

import(
	"fmt"
)

func main() {
	var a, b, c, d int64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	if a <= b + d +c{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")

	}
}
