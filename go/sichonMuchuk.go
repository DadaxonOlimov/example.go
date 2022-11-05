package main

import (
	"fmt"
	"math"
)

func main() {
	var cat1, cat2, mouse int
	fmt.Scan(&cat1, &cat2, &mouse) 
	if math.Abs(float64(mouse) - float64(cat1)) > math.Abs(float64(mouse) - float64(cat2)){
		fmt.Println("2-mushuk")
	}else if  math.Abs(float64(mouse) - float64(cat1)) < math.Abs(float64(mouse) - float64(cat2)){
		fmt.Println("1-mushuk")
	}else{
		fmt.Printf("%s\n","sichqon")
	}	
	
}