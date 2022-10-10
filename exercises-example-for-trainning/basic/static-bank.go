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

func SeeBalance(option int) {

	OptionsValidation(option)

	var balance = Account{"3876", 300.02, "Bank Hash Deck"}
	fmt.Println("\tBalance Information: ")
	fmt.Println("\tAccount Number: ", balance.AccountNumber)
	fmt.Println("\tAccount Balance: ", balance.Balance)
	fmt.Println("\tAccount Bank: ", balance.Bank)
}

func WithdrawMoney(option int) {
	OptionsValidation(option)

	fmt.Println("Inform the withdrawal amount $: ")

	var value float64
	fmt.Scanln(&value)

	var account = Account{"3876", 300.02, "Bank Hash Deck"}

	var withdrawal = account.Balance - value
	account.Balance = withdrawal

	fmt.Println("\t Withdrawal was successful!  ")
	fmt.Println("\t Balance Information: ")
	fmt.Println("\tAccount Number: ", account.AccountNumber)
	fmt.Println("\tAccount Balance: ", account.Balance)
	fmt.Println("\tAccount Bank: ", account.Bank)
}

func TransferMoney(option int) {
	OptionsValidation(option)

	var account = Account{"3876", 800.02, "Bank Hash Deck"}

	fmt.Println("Inform destination account number: ")
	var destinationAccountNumber string
	fmt.Scanln(&destinationAccountNumber)

	fmt.Println("Inform the value for transfer: ")
	var transferValue float64
	fmt.Scanln(&transferValue)

	var value = account.Balance - transferValue
	account.Balance = value

	fmt.Println("\t Transfer was successful!  ")
	fmt.Println("\t Account Destination Information: ")
	fmt.Println("\tAccount Number: ", destinationAccountNumber)
	fmt.Println("\tAccount Bank: ", "American Digital Bank")
	fmt.Println("\tAmount transferred: ", transferValue)

	fmt.Println("\t Your balance: ")
	fmt.Println("\tAccount Number: ", account.AccountNumber)
	fmt.Println("\tAccount Balance: ", account.Balance)
	fmt.Println("\tAccount Bank: ", account.Bank)
}

func DepositMoney(option int) {
	OptionsValidation(option)
}

func ReceiveOptions(options int) {

	switch options {
	case 1:
		SeeBalance(options)
	case 2:
		WithdrawMoney(options)
	case 3:
		TransferMoney(options)
	case 4:
		DepositMoney(options)
	default:
		fmt.Println("Options not found")
	}
}

func OptionsValidation(options int) {

	if options == 0 {
		fmt.Println(options, "invalid option!")
	}
}

func main() {
	fmt.Println("What do you search? Enter in a options: ")
	fmt.Println("1.\tSee balance \n2.\tWithdraw Money \n3.\tTransfer \n4.\tDeposit")
	var options int
	fmt.Scanln(&options)
	ReceiveOptions(options)
}
