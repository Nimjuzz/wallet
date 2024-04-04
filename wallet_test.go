package wallet

import (
    "testing"
)

func TestWallet(t *testing.T) {
    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw with sufficient balance", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(10))
        assertNoError(t, err)
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw with insufficient balance", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(30))
        assertError(t, err, "insufficient balance")
        assertBalance(t, wallet, Bitcoin(20)) // Balance should remain unchanged
    })
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
    t.Helper()
    got := wallet.Balance()
    if got != expected {
        t.Errorf("got %f BTC, want %f BTC", got, expected)
    }
}

func assertError(t *testing.T, err error, expectedMsg string) {
    t.Helper()
    if err == nil {
        t.Errorf("expected an error but got none")
    }
    if err.Error() != expectedMsg {
        t.Errorf("got error %q, want %q", err.Error(), expectedMsg)
    }
}

func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
}