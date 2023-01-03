package main

// The actual value currency 2022-01-03 8am

import "fmt"

type Currency struct {
	USD string
	EUR string
	BRL string
	AUD string
	CAD string
}

type CurrencyValue struct {
	Currency string
	Value    float64
	Country  string
}

func main() {

	fmt.Println("Select your actual currency: ")
	fmt.Println("1.\tUSD \n2.\tEUR \n3.\tBRL \n4.\tAUD \n5.\tCAD")

	var actualCurrency int
	fmt.Scanln(&actualCurrency)

	fmt.Println("Select the currency to do the conversion : ")
	fmt.Println("1.\tSee balance \n2.\tWithdraw Money \n3.\tTransfer \n4.\tDeposit")

	var currencyConvert int
	fmt.Scanln(&currencyConvert)

	CurrencyConvert(actualCurrency, currencyConvert)

}

func CurrencyConvert(option int, actualCurrency int) {

	switch option {
	case 1:
		ConvertToDolar(actualCurrency)
	case 2:
		ConvertToEuro(actualCurrency)
	case 3:
		ConvertToBRL(actualCurrency)
	case 4:
		ConvertToAUD(actualCurrency)
	case 5:
		ConvertToCAD(actualCurrency)
	default:
		fmt.Println("Options not found")
	}

}

func ConvertToDolar(actualCurrency int) {
// to calculate
}

func ConvertToEuro(actualCurrency int) {
// to calculate
}

func ConvertToBRL(actualCurrency int) {
// to calculate
}

func ConvertToAUD(actualCurrency int) {
// to calculate
}

func ConvertToCAD(actualCurrency int) {
// to calculate
}
