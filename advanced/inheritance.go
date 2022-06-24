package main

import "fmt"

type ferrari struct {
	Name       string
	PowerTurbo bool
	ActVelocity int
}

func main() {
	f := ferrari{}
	f.Name = "F40"
	f.ActVelocity = 0
	f.PowerTurbo = true

	fmt.Println("The ferrari is turbo charged? %v\n", f.Name, f.PowerTurbo)
}
