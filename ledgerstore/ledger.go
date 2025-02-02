package ledgerstore

import (
	"errors"
	"sync"
)

// LedgerStore structure to hold transactions
type LedgerStore struct {
	mu           sync.Mutex
	balance      int
	transactions []TransactionStore
}

// TransactionStore structure
type TransactionStore struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

// NewLedgerStore creates a new instance
func NewLedgerStore() *LedgerStore {
	return &LedgerStore{
		balance:      0,
		transactions: []TransactionStore{},
	}
}

// Deposit money
func (l *LedgerStore) Deposit(amount int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.balance += amount
	l.transactions = append(l.transactions, TransactionStore{Amount: amount, Type: "deposit"})
}

// Withdraw money
func (l *LedgerStore) Withdraw(amount int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.balance < amount {
		return errors.New("insufficient funds")
	}

	l.balance -= amount
	l.transactions = append(l.transactions, TransactionStore{Amount: amount, Type: "withdrawal"})
	return nil
}

// GetBalance returns current balance
func (l *LedgerStore) GetBalance() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.balance
}

// GetTransactions returns transaction history
func (l *LedgerStore) GetTransactions() []TransactionStore {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.transactions
}
