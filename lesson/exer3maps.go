package main

import (
	"fmt"
	"log"
)

func main()  {
	var m int
	fmt.Println("Please enter num: ")
	_,err := fmt.Scanln(&m)
	if err !=nil{
		log.Fatal("Error occured: ", err)
		
	}
	fmt.Printf("%v", maps(m))
}


func maps(m int)map[int]int { 
	number :=make(map[int]int,m)
	for i :=1; i<m; i++{
		number[i] = i *i
	}
	return number

}