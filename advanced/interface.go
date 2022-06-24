package main

import "fmt"

type print interface {
	toString() string
}

type person struct {
	name     string
	lastName string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - $ %.2f", p.name, p.price)
}

func ToPrint(i print) {
	fmt.Println(i.toString())
}

func main() {
	var printPerson print = person{"Jhon", "Snow"}
	fmt.Println(printPerson.toString())
	ToPrint(printPerson)

	var printProduct = product{"Smartphone", 500}
	fmt.Println(printPerson.toString())
	ToPrint(printProduct)
}
