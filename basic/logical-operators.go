package main

import "fmt"

// return bool variable
func Shopping(freela1, freela2 bool) (bool, bool, bool) {
	shoppingTv50 := freela1 && freela2
	shoppingTv32 := freela1 != freela2
	shoppingIceCream := freela1 || freela2

	return shoppingTv50, shoppingTv32, shoppingIceCream
}

func main() {
	tv50, tv32, sorvete := Shopping(true, true)
	fmt.Printf("TV: %t, TV32: %t, Ice-Cream: %t", tv50, tv32, sorvete, !sorvete)
}
