package main

import "fmt"

func GetAprovedValue(number int) int {
	fmt.Println("End!")

	if number > 5000 {
		fmt.Println("value high...")
		return 5000
	} else {
		fmt.Println("Value low...")
		return number
	}
}

func main() {
	fmt.Println(GetAprovedValue(2000))
}
