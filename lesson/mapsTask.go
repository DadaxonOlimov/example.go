package main

import "fmt"

func main()  {
	fmt.Println(array())

}


func array()int  {
	var n int
	var a int
	fmt.Scan(&n)
	var sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += a
	}
	return sum 
}
