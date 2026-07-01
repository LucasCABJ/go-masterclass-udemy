package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance float64
	OwnerName string
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	acc.Balance += amount
	fmt.Printf("Deposited %.2f to account %s. New balance: %.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be positive")
	}
	if amount > acc.Balance {
		return errors.New("Insufficient funds")
	}
	acc.Balance -= amount
	fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account Number: %s, Owner: %s, Balance: %.2f", acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account
	InterestRate float64 // e.g., 0.02 for 2%
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest %.2f to account %s\n", interest, sa.AccountNumber)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Println("Error adding interest:", err)
	}
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdraw amount must be positive")
	}
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return errors.New("Insufficient funds, overdraft limit exceeded")
	}
	oa.Balance -= amount
	fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {
	savings := SavingsAccount{
		Account: Account{
			AccountNumber: "SA123",
			Balance:       1000.0,
			OwnerName:     "Alice",
		},
		InterestRate: 0.02,
	}

	overdraft := OverdraftAccount{
		Account: Account{
			AccountNumber: "OA456",
			Balance:       500.0,
			OwnerName:     "Bob",
		},
		OverdraftLimit: 200.0,
	}

	fmt.Println(savings)
	fmt.Println(overdraft)

	savings.Deposit(200)
	savings.Withdraw(50)
	savings.AddInterest()

	overdraft.Withdraw(600) // Should succeed due to overdraft
	err := overdraft.Withdraw(200) // Should fail due to overdraft limit
	if err != nil {
		fmt.Println("Error:", err)
	}
}