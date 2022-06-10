package main

import "fmt"

func main() {
	var shirtPrice [3]float64
	fmt.Println(shirtPrice)

	shirtPrice[0], shirtPrice[1], shirtPrice[2] = 50, 42.20, 32.14

	fmt.Println(shirtPrice)

	total := 0.0
	for i := 0; i < len(shirtPrice); i++ {
		total += shirtPrice[i]
	}

	fmt.Println("Total Shirt prices ", total)

}
