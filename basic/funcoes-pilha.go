package main

import "runtime/debug"

func functionThree() {
	debug.PrintStack()
}

func functionTwo() {
	functionThree()
}

func functionOne() {
	functionTwo()
}

func main() {
	functionOne()
}
