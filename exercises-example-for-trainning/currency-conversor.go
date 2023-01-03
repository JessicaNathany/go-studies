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

	fmt.Println("What value do you want to convert: ")

	var convertValue float
	fmt.Scanln(&convertValue)

	fmt.Println("Select the currency to do the conversion : ")
	fmt.Println("1.\tSee balance \n2.\tWithdraw Money \n3.\tTransfer \n4.\tDeposit")

	var currencyOptions int
	fmt.Scanln(&currencyOptions)

	CurrencyConvert(convertValue, currencyOptions)
}

func CurrencyConvert(option int, convertValue int) {

	switch option {
	case 1:
		ConvertToDolar(convertValue)
	case 2:
		ConvertToEuro(convertValue)
	case 3:
		ConvertToBRL(convertValue)
	case 4:
		ConvertToAUD(convertValue)
	case 5:
		ConvertToCAD(convertValue)
	default:
		fmt.Println("Options not found")
	}

}

func ConvertToDolar(convertValue int) {

	var dolar = CurrencyValue{"USD", 5.36, "EUA"}
	var conversionValue = convertValue / int(dolar.Value)

	fmt.Println("\tCom ", convertValue)
	fmt.Println("\tYou can buy ", conversionValue)

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
