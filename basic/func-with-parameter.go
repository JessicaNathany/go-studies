package main

import "fmt"

func Multiplication(a, b int) int {
	return a * b
}

// func received two parameters int and return paramter int
func Exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	result := Exec(Multiplication, 3, 4)

	fmt.Println(result)
}
