package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	a2 := []int{1, 2, 3}  // slice
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(a2))

	b1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b1)

	// slice it no a array! Slice definy a peace the array
	c1 := b1[1:3]
	fmt.Println(c1, b1)

	c3 := b1[:2]
	fmt.Println(b1, c3)
}
