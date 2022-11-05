package main

import "fmt"


func main()  {
	var a,b int
	fmt.Scan(&a,&b)
	fmt.Println(hisobla(a,b))
}

func hisobla(athletes int, failed int )int {
	var nextRound int
	var massiv []int
	for i := 0; i < athletes; i++ {
		fmt.Scan(&nextRound)
		massiv = append(massiv, nextRound)
	} 
	for i := 0; i < athletes; i++ {
		for j := 0;  j < athletes; j++ {
			if massiv[i] > massiv[j] {
				massiv[i],massiv[j] = massiv[j],massiv[i]
			}
		}
	}
	num := massiv[failed]
	for i, v := range massiv {
		if v < num {
			return i
		}
	}
	return 0
}

// 8 5
// 10 9 8 7 7 7 5 5