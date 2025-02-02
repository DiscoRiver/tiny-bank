package ledgerstore

import (
	"errors"
	"github.com/discoriver/tiny-bank/model"
	"sync"
)

// LedgerStore structure to hold transactions
type LedgerStore struct {
	mu           sync.Mutex
	balance      model.Balance
	transactions model.Transactions
}

// NewLedgerStore creates a new instance
func NewLedgerStore() *LedgerStore {
	return &LedgerStore{
		balance:      model.Balance{Amount: 0},
		transactions: model.Transactions{},
	}
}

// Deposit money
func (l *LedgerStore) Deposit(amount int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.balance.Amount += amount
	l.transactions.Transactions = append(l.transactions.Transactions, model.Transaction{Amount: amount, Type: "deposit"})
}

// Withdraw money
func (l *LedgerStore) Withdraw(amount int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.balance.Amount < amount {
		return errors.New("insufficient funds")
	}

	l.balance.Amount -= amount
	l.transactions.Transactions = append(l.transactions.Transactions, model.Transaction{Amount: amount, Type: "withdrawal"})
	return nil
}

// GetBalance returns current balance
func (l *LedgerStore) GetBalance() model.Balance {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.balance
}

// GetTransactions returns transaction history
func (l *LedgerStore) GetTransactions() model.Transactions {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.transactions
}
