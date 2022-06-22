package main

func Increment(n int) {
	n++ // same n = n + 1
}

func Increment2(n *int) {
	// revisão: é usado para desreferenciar, ou seja,
	//ter acesso ao valor no qual ponteiro aponta
	*n++
}

func main() {
	n := 1

	Increment(n)
}
