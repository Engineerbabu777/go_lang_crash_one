package main

import "fmt"

// Account represents a bank account with an ID and balance.
type Account struct {
	ID      int
	Balance float64
}

// CreateAccount creates a new account with the specified ID and initial balance.
func CreateAccount(accountID int, initialBalance float64) *Account {
	return &Account{ID: accountID, Balance: initialBalance}
}

// Deposit adds funds to the account.
func Deposit(account *Account, amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid deposit amount.")
		return
	}

	account.Balance += amount
	fmt.Printf("Deposit of %.2f successful. New balance: %.2f\n", amount, account.Balance)
}

// Withdraw deducts funds from the account.
func Withdraw(account *Account, amount float64) {
	if amount <= 0 || amount > account.Balance {
		fmt.Println("Invalid withdrawal amount.")
		return
	}

	account.Balance -= amount
	fmt.Printf("Withdrawal of %.2f successful. New balance: %.2f\n", amount, account.Balance)
}

// CheckBalance returns the balance of the account.
func CheckBalance(account *Account) float64 {
	return account.Balance
}

func main() {
	// Create a new account with ID 1 and initial balance of 100.00
	account := CreateAccount(1, 100.00)

	// Perform some transactions
	Deposit(account, 70.00)
	// Check the final balance
	finalBalance := CheckBalance(account)
	fmt.Printf("Final balance: %f\n", finalBalance)
	Withdraw(account, 30.00)

	// Check the final balance
	// finalBalance := CheckBalance(account)
	// fmt.Printf("Final balance: %.2f\n", finalBalance)
}
