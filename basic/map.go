package main

import "fmt"

func main() {

	approved := make(map[int]string)

	approved[90768610052] = "Maria"
	approved[38637616003] = "Jéssica"
	approved[30358236002] = "Jhow"

	fmt.Println(approved)

	for document, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, document)
	}

	fmt.Println(approved[38637616003])
	delete(approved, 30358236002)
	fmt.Println(approved[90768610052])
}
