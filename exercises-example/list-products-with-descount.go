package main

import "fmt"

// Eletronic products has 10% the discount and Sports products has 5% the discount
type Product struct {
	Name     string
	Type     int
	discount float64
	Price    float64
}

func PrintProduct(productType int) {

	if productType == 1 {
		var sportsProducts = Product{"Bascket Ball", 1, 0.05, 100}
		fmt.Println(sportsProducts, sportsProducts.CalculatePriceWithFiveDiscountPercentage())
	} else {
		var eletronicProduct = Product{"Macbook", 2, 0.10, 7000}
		fmt.Println(eletronicProduct, eletronicProduct.CalculatePriceWithTemDiscountPercentage())
	}
}

func (product Product) CalculatePriceWithFiveDiscountPercentage() float64 {
	descountPercentage := product.Price * product.discount / 100
	newPrice := product.Price - descountPercentage

	return newPrice
}

func (product Product) CalculatePriceWithTemDiscountPercentage() float64 {
	descountPercentage := product.Price * 10 / 100
	newPrice := product.Price - descountPercentage

	return newPrice
}

func ReceiveProducts(productType int) {

	switch productType {
	case 1:
		PrintProduct(1)
	case 2:
		PrintProduct(2)
	default:
		fmt.Println("Not found")
	}
}

func main() {
	fmt.Println("What do you search? Enter in a options: ")
	fmt.Println("1.\tSports Products\n2.\tEletronic Products")
	var options int
	fmt.Scanln(&options)

	ReceiveProducts(options)
}
