package main

import "fmt"

// Creating and add vector

func main() {
	var products [5]string

	for i := 0; i < len(products); i++ {
		fmt.Print("nform the product name: ")
		fmt.Scanf("%s", &products[i])
	}

	fmt.Println("Your products are: ")
	for _, product := range products {
		fmt.Printf("_%s\n", product)
	}

	// start array

	var matriz [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Inform a number: ")
			fmt.Scanf("%d", &matriz[i][j])
		}
	}
	fmt.Println(matriz)
}
