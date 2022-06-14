package main

import "fmt"

func main() {
	funcForLetter := map[string]map[string]float64{
		"B": {
			"Bruce Banner": 15456.78,
			"Bruce Wayne":  8456.78,
		},
		"J": {
			"Jessy Pinckman": 11325.45,
		},
		"T": {
			"Tony Startk": 1400.0,
		},
	}

	delete(funcForLetter, "J") // delete empolyer with leter J

	for letter, funcs := range funcForLetter {
		fmt.Println(letter, funcs)
	}
}
