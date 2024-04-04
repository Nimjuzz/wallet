package wallet

import (
    "errors"
    "sync"
)

// Bitcoin represents the amount of bitcoins in the wallet.
type Bitcoin float64

// Wallet represents a Bitcoin wallet with deposit, withdraw, and balance operations.
type Wallet struct {
    balance Bitcoin
    mutex   sync.Mutex
}

// Deposit adds a specified amount of bitcoins to the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
    w.mutex.Lock()
    defer w.mutex.Unlock()
    w.balance += amount
}

// Withdraw subtracts a specified amount of bitcoins from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
    w.mutex.Lock()
    defer w.mutex.Unlock()
    if amount > w.balance {
        return errors.New("insufficient balance")
    }
    w.balance -= amount
    return nil
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
    w.mutex.Lock()
    defer w.mutex.Unlock()
    return w.balance
}