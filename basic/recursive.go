package main

import "fmt"

// method return error
func Factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		FactorialBefore, _ := Factorial(n - 1)
		return n * FactorialBefore, nil
	}
}

func main() {
	result, _ := Factorial(5)
	fmt.Println(result)

	//_, => ignore result
	_, err := Factorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
