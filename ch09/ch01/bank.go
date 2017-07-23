// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)
var withdrawResult = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawResult
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if balance-amount >= 0 {
				balance -= amount
				withdrawResult <- true
			} else {
				withdrawResult <- false
			}
		}
	}
}

func init() {
	go teller()
}
