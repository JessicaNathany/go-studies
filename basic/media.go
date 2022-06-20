package main

import "fmt"

func Media(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Println("Média: %.2f", Media)
}
