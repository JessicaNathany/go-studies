package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum => ", a+b)
	fmt.Println("Subtraction => ", a-b)
	fmt.Println("Division => ", a/b)
	fmt.Println("Multiplication => ", a*b)
	fmt.Println("Module => ", a%b)

	// bitwise
	fmt.Println("E =>", a&b)    // 11 & 10 = 10
	fmt.Println("Ou => ", a|b)  // 11 | 10 = 11
	fmt.Println("Xor => ", a^b) // 11 ^10 = 01

	//c := 3.0
	//b := 2.0

	// tohers operation using math
	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))

}
