package main

import (
	"fmt"
	"strings"
)

type person struct {
	Name     string
	LastName string
}

func (p person) GetFullName() string {
	return p.Name + " " + p.LastName
}

func (p *person) SetFullName(fullNmae string) {
	parts := strings.Split(fullNmae, " ")
	p.Name = parts[0]
	p.LastName = parts[1]
}

func main() {
	p1 := person{"Jéssica", "Carvalho"}
	fmt.Println(p1.GetFullName())

	p1.SetFullName("Thiago Carvalho")
	fmt.Println(p1.GetFullName())
}
