package main

import "fmt"

func main() {
	var rows int
	fmt.Scan(&rows)
	fmt.Println(draw(rows))
}

func draw(countOfRows int) string {
	initialStr := fmt.Sprintf("1\n")
	counter := 2
	for i := 1; i < countOfRows; i++ {
		for j := 0; j <= i; j++ {
			initialStr += fmt.Sprintf("%d ", counter)
			counter++
		}
		initialStr += fmt.Sprintf("\n")
	}

	return initialStr
}
