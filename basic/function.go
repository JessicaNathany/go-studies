package main

import "fmt"

func funcTakesNoParametersAndHasNoReturn() {
	fmt.Println("Funciton 1")
}

func funcWithParameters(p1 string, p2 string) {
	fmt.Printf("Function 2: %s %s\n", p1, p2)
}

func funcWithReturn() string {
	return "Function 3"
}

func funcTakesTwoParametersAndHasReturn(p1, p2 string) string {
	return fmt.Sprintf("Function 4: %s %s", p1, p2)
}

func funcHasTwoReturns() (string, string) {
	return "Return 1", "Return 2"
}

func main() {
	funcTakesNoParametersAndHasNoReturn()
	funcWithParameters("param1", "param2")
	return3, return4 := funcWithReturn(), funcTakesTwoParametersAndHasReturn("Param 1", "Param 2")

	fmt.Println(return3)
	fmt.Println(return4)

	return51, return52 := funcHasTwoReturns()

	fmt.Println("Function 5:", return51, return52)
}
