package main

import (
	"errors"
	"fmt"
	"log"	
	_ "os"	//this is ignored 
)

// this defines a new type
type BankAC struct {
	acno int
	balance int
}

func (m *BankAC) getCurrentBalance() int { 
	return m.balance 
}

func (m *BankAC) deposit(amount int) {
	m.balance += amount
}

func (m *BankAC) withdraw(amount int) error {
	if amount < 0 { 
		return errors.New("Amount should be positive")
	} else if m.balance < amount {
		return errors.New("Insufficient balance")
	} 	

	m.balance -= amount
	return nil
}

func main() {
	ac1 := &BankAC{acno:1, balance: 1000}	

	fmt.Println(ac1.getCurrentBalance())	
	ac1.deposit(1000)
	fmt.Println(ac1.getCurrentBalance())
	res := ac1.withdraw(500)
	if res != nil {
		log.Println(res)
	} else {
		fmt.Println("Amount Withdrawn Successfully")
	}
	fmt.Println(ac1.getCurrentBalance())

}