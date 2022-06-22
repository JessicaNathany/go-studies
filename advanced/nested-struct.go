package main

import "fmt"

type item struct {
	productId int
	quantity  int
	price     float64
}

type order struct {
	userId int
	items  []item
}

func (o order) TotalValue() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	order := order{
		userId: 1,
		items: []item{
			item{productId: 1, quantity: 2, price: 12.10},
			item{productId: 2, quantity: 1, price: 23.50},
			item{productId: 11, quantity: 100, price: 3.13},
		},
	}

	fmt.Println("The value order is $ %.2f", order.TotalValue())
}
