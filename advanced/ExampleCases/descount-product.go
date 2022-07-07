package main

// The program read the product price and return the new price with 5% descount 10*5/100 = 0.5 | 1500*10/100

import (
	"fmt"
)

type Person struct {
	Name               string
	Salary             float64
	DescountPercentage int
}

func ProductCalculate(productPrice float64) float64 {

	descountPercentage := productPrice * 5 / 100
	newPrice := productPrice - descountPercentage

	return newPrice
}

func main() {

	fmt.Println("What is the price of product?")
	var productPrice float64
	fmt.Scanln(&productPrice)

	newPriceProduct := ProductCalculate(productPrice)
	fmt.Printf("Hi, the price of product is: %.2f", newPriceProduct)
}
