package main

import "fmt"

func PrintListAproved(aproveds ...string) {
	fmt.Println("List aproved")

	for i, aproved := range aproveds {
		fmt.Printf("%d) %s\n", i+1, aproved)
	}
}

func main() {
	aproveds := []string{"Mary", "Jhon", "Jéssica"}
	PrintListAproved(aproveds...)
}
