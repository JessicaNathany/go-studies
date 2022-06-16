package main

import "fmt"

func change(p1, p2 int) (second int, first int) {
	second = p2
	first = p1
	return // clear return
}

func main() {
	x, y := change(2, 3)
	fmt.Println(x, y)

	second, first := change(7, 1)
	fmt.Println(second, first)
}
