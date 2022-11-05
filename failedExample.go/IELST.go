package main

import (
	"fmt"
	"math"
)

func main()  {
	var speaking, reading []float64
	var listening, writing float64
	speaking = GetScores()
	reading = GetScores()
	fmt.Scan(&listening, &writing)
	result := Ielst(speaking,reading,listening,writing)
	fmt.Println(result)
}

func Ielst(speaking, reading []float64, listening, writing float64) float64 {
	speak, read := PlusFunk(speaking), PlusFunk(reading)
	return math.Ceil((speak + read + listening + writing) / 4)
}

func GetScores() []float64 {
	var slice []float64
	var a float64
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		slice = append(slice, a)
		
	}
	return slice
}

func PlusFunk(scores []float64) float64 {
	var sum float64
	for _, v := range scores {
		sum += v
	}
	return math.Floor(sum / 4)
}