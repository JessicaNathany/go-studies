package main

import "fmt"

func main() {
	employerSalary := map[string]float64{
		"Jéssica Nathany": 11325.45,
		"Gabriela Silva":  15456.78,
		"MAria Marruá":    1200.0,
	}

	employerSalary["rafel Filho"] = 1350.0
	delete(employerSalary, "Undefined") // try deletes an element that does not exist

	for name, salary := range employerSalary {
		fmt.Println(name, salary)
	}
}
