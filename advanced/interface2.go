package main

import "fmt"

type sportCar interface {
	turboPower()
}

type Ferrari struct {
	model       string
	turboOn     bool
	Actvelocity int
}

func (f *Ferrari) turboPower() {
	f.turboOn = true
}

func main() {
	car1 := Ferrari{"F40", false, 0}
	car1.turboPower()

	var car2 sportCar = &Ferrari{"F40", false, 0}
	car2.turboPower()

	fmt.Println(car1, car2)
}
