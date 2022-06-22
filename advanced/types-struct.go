package main

import "fmt"

type Product struct {
	name     string
	price    float64
	descount float64
}

type Address struct {
	Name    string
	City    string
	Country string
}

// Method with receptor

func (p Product) priceDescount() float64 {
	return p.price * (1 - p.descount)
}

func (p Address) GetCoordinates() string {
	return "23°34'05.9 S 46°41'28.4 W"
}

func main() {

	var product1 Product
	product1 = Product{
		name:     "Skate",
		price:    150,
		descount: 0.05,
	}

	var address Address
	address = Address{
		Name:    "Jéssica",
		City:    "São Paulo",
		Country: "Brazil",
	}

	fmt.Println(address, address.GetCoordinates())
	fmt.Println(product1, product1.priceDescount())

	product2 := Product{"Notebook", 1989.90, 0.10}
	fmt.Println(product2, product2.priceDescount())
}
