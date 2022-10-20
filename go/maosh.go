package main

import(
	"fmt"
)

func main(){
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	if a > b&& b>c || a > c&& c>b{
		fmt.Println(a)
	}else if b > a&& a>c || b > c&& c>a{
		fmt.Println(a)
	}else{
		fmt.Println(c)
	}
}

