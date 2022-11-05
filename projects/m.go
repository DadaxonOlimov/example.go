package main

import (
	"fmt"
	"strings"
)

 func main()  {
	var str string
	fmt.Scan(&str)

	fmt.Println(mapss(str))

}

func mapss(s string)map[string]int{
	a := make(map[string]int)		
		Slise := strings.Split(s,"")
		for _,i := range Slise{
			if _,ok := a[i]; ok{
				a[i] += 1
		}else{
			a[i] = 1
		}
	}
	return a
}

