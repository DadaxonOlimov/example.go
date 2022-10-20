package main

import "fmt"

func main()  {
	// var a int
	// fmt.Scan(&a)
	// if 2 < a && a < 6 {
	// 	fmt.Println("Spring")
	// }else if 5 < a && a < 9 {
	// 	fmt.Println("Summer")
	// }else if 8 < a && a <12 {
	// 	fmt.Println("Automn")
	// }else if a < 3 || a == 12 {
	// 	fmt.Println("Winter")
	// }
	 a := map[string]bool{"fARHOD" : true}
	 fmt.Println(a["asd"])

	delete(a, "fARHOD")
	_, ok := a["fARHOD"]

	fmt.Println(a)
}


