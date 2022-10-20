package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a % 400 ==0 {
		fmt.Println("Kabisa yili")
	}else if  a % 4 ==0 && a % 100 !=0 {
		fmt.Println("Kabisa yili")
	}else{
		fmt.Println("Kabisa yili emas")
	}
}
	

