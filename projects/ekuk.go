package main

import "fmt"

func main() {

}

func ekuk(int, int, int) string {
	var a int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Println(a%2 == 0)
	}
	return "go"
}
