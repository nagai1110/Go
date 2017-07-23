// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package bank

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	Deposit(200)

	ok := Withdraw(100)
	if !ok {
		t.Errorf("Withdraw(100) = false, want true")
	}

	ok = Withdraw(100)
	if !ok {
		t.Errorf("Withdraw(100) = false, want true")
	}

	ok = Withdraw(100)
	if ok {
		t.Errorf("Withdraw(100) = true, want false")
	}
}
