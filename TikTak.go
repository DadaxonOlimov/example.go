package main

import "fmt"

func main()  {
	// var x string
	// fmt.Scan(&x)
	var o string
	var massiv []string
	for i := 0; i < 3; i++ {
			fmt.Scan(&o)
			massiv = append(massiv,o)
		for _, v := range massiv {
			if massiv[v] == "x"		
		}
		}
}

// func TikTak()  {
// 	var x string
// 	fmt.Scan(&x)
// 	var o string
// 	var massiv []string
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			fmt.Scan(&o)
// 			massiv = append(massiv,o)
// 		}
// 	}
// }