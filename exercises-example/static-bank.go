package main

import "fmt"

type Person struct {
	Name  string
	Email string
}

type Account struct {
	AccountNumber string
	Balance       float64
	Bank          string
}

func WithdrawMoney(option int) {
	// option 2
}

func SeeBalance(option int) {
	// option 1
}

func TransferMoney(option int) {
	// option 3
}

func DepositMoney(option int) {
	// option 4
}

func ReceiveOptions(options int) {

	switch options {
	case 1:
		SeeBalance(1)
	case 2:
		WithdrawMoney(2)
	case 3:
		TransferMoney(2)
	case 4:
		DepositMoney(2)
	default:
		fmt.Println("Not found")
	}
}

func main() {
	fmt.Println("What do you search? Enter in a options: ")
	fmt.Println("1.\tSee balance \n2.\tWithdraw Money \n3.\tTransfer \n4.\tDeposit")
	var options int
	fmt.Scanln(&options)
	ReceiveOptions(options)
}
