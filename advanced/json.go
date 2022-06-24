package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	product1 := product{1, "Notebook", 1899.90, []string{"Promotion", "Eletronic"}}
	productJson, _ := json.Marshal(product1)
	fmt.Println(string(productJson))

	// json to struct
	var product2 product
	jsonString := `{"id":2, "name":"pen", "price":8.90, "tags":["Paper", "Import"]}`
	json.Unmarshal([]byte(jsonString), &product2)
	fmt.Println(product2.Tags[1])
}
