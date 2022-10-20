package main

import "fmt"

func main()  {
	fmt.Println(array())

}

func array() []int  {
	var a []int
	var b int
	var str int
	fmt.Scan(&b)

	for i := 0; i < b; i++ {
		fmt.Scan(&str)
		a = append(a, str)
	}

	var sum int
	for _, v := range a {
		sum += v
	}

	return a
}

// // 5
// // *#  f 
// // *#
// // #*
// // *#
// // *#

