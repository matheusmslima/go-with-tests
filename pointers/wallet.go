package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error {

	if amount > wallet.balance {
		return ErrInsufficientFunds
	}

	wallet.balance -= amount
	return nil
}

func ProcessWithdrawal(wallet *Wallet, accountID string, amount Bitcoin) error {
	if err := wallet.Withdraw(amount); err != nil {
		return fmt.Errorf("processing withdrawal for account %s: %w", accountID, err)
	}
	return nil
}
