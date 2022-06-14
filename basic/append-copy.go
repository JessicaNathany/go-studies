package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 1, 2, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	//make a copy slice1 for slice2
	copy(slice2, slice1) // slice2 receive the value an copy the slice1
	//slice2 => destination slice2 => origin
	fmt.Println(slice2)
}
