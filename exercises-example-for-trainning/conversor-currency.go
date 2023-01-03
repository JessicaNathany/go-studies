package main

// The actual value currency 2022-01-03 8am

import "fmt"

type Currency struct {
	USD string
	EUR string
	AUD string
	CAD string
}

type CurrencyValue struct {
	Currency string
	Value    float64
	Country  string
}

func main() {

	fmt.Println("What value do you want to convert: ")

	var convertValue float64
	fmt.Scanln(&convertValue)

	fmt.Println("Select the currency to do the conversion : ")
	fmt.Println("1.\tDolar \n2.\tEuro \n3.\tAUD \n4.\tCAD")

	var currencyOptions int
	fmt.Scanln(&currencyOptions)

	CurrencyConvert(currencyOptions, convertValue)
}

func CurrencyConvert(option int, convertValue float64) {

	switch option {
	case 1:
		ConvertToDolar(convertValue)
	case 2:
		ConvertToEuro(convertValue)
	case 3:
		ConvertToAUD(convertValue)
	case 4:
		ConvertToCAD(convertValue)
	default:
		fmt.Println("Options not found")
	}

}

func ConvertToDolar(convertValue float64) {

	var dolar = CurrencyValue{"USD", 5.36, "EUA"}
	var total = convertValue / dolar.Value

	fmt.Println("\tCom ", convertValue)
	fmt.Println("\tYou can buy ", total)
}

func ConvertToEuro(convertValue float64) {
	var euro = CurrencyValue{ "EUR", 5.77, "UE" }
	var total = convertValue / euro.Value

	fmt.Println("\tCom ", convertValue)
	fmt.Println("\tYou can buy ", total)
}

func ConvertToAUD(convertValue float64) {
	var au = CurrencyValue{ "AUD", 3.68, "AU" }
	var total = convertValue / au.Value

	fmt.Println("\tCom ", convertValue)
	fmt.Println("\tYou can buy ", total)
}

func ConvertToCAD(convertValue float64) {
	var cad = CurrencyValue{ "CAD", 3.68, "CA" }
	var total = convertValue / cad.Value

	fmt.Println("\tCom ", convertValue)
	fmt.Println("\tYou can buy ", total)
}
