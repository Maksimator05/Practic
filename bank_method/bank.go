package main

import "fmt"

type BankAccount struct {
	Balance float64
}

func (acc *BankAccount) Deposit(amount float64) {
	acc.Balance += amount
}
func (acc *BankAccount) Withdraw(amount float64) bool {
	if amount > acc.Balance {
		return false
	}
	acc.Balance -= amount
	return true
}
func main() {

	account := BankAccount{1000}

	account.Deposit(500)
	fmt.Println("Balance after push", account.Balance)

	success := account.Withdraw(500)
	if success {
		fmt.Println("Take success")
	} else {
		fmt.Println("Take failure")
	}
	fmt.Println("Current balance", account.Balance)
}
