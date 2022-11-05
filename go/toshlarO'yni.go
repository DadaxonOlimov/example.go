package main

import "fmt"

func main()  {
	var Stone int
	fmt.Scan(&Stone)
	if Stone % 2 == 1 {
		fmt.Println("First player")
	}else if Stone % 2 ==0{
		fmt.Println("Second player")
	}
}	