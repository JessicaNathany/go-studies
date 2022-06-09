package main

import "fmt"

func GetCriptoCurrencyValue(criptoName string) float64 {
	var cripto = (criptoName)

	switch cripto {
	case "BTC":
		return 30.11559
	case "ETH":
		return 1.79376
	case "USDC":
		return 1.00
	case "BNB":
		return 289.42
	case "XRP":
		return 0.3984
	default:
		return 0
	}
}

func main() {
	fmt.Println("BTC", GetCriptoCurrencyValue("BTC"))
	fmt.Println("ETH",GetCriptoCurrencyValue("ETH"))
	fmt.Println("USDC",GetCriptoCurrencyValue("USDC"))
	fmt.Println("BNB",GetCriptoCurrencyValue("BNB"))
	fmt.Println("XRP",GetCriptoCurrencyValue("XRP"))
}
