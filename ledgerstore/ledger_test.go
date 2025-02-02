package ledgerstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLedger(t *testing.T) {
	l := NewLedgerStore()
	assert.NotNil(t, l, "LedgerStore should be initialized")
	assert.Equal(t, 0, l.GetBalance().Amount, "Initial balance should be 0")
	assert.Empty(t, l.GetTransactions(), "TransactionStore history should be empty")
}

func TestDeposit(t *testing.T) {
	l := NewLedgerStore()

	l.Deposit(100)
	assert.Equal(t, 100, l.GetBalance().Amount, "Balance should be updated after deposit")

	l.Deposit(50)
	assert.Equal(t, 150, l.GetBalance().Amount, "Balance should reflect multiple deposits")

	transactions := l.GetTransactions()
	assert.Len(t, transactions.Transactions, 2, "TransactionStore count should match deposits")
	assert.Equal(t, "deposit", transactions.Transactions[0].Type, "TransactionStore type should be deposit")
}

func TestWithdraw_Success(t *testing.T) {
	l := NewLedgerStore()
	l.Deposit(200)

	err := l.Withdraw(100)
	assert.NoError(t, err, "Withdraw should succeed if funds are available")
	assert.Equal(t, 100, l.GetBalance().Amount, "Balance should be updated after withdrawal")

	transactions := l.GetTransactions()
	assert.Len(t, transactions.Transactions, 2, "TransactionStore count should include withdrawals")
	assert.Equal(t, "withdrawal", transactions.Transactions[1].Type, "TransactionStore type should be withdrawal")
}

func TestWithdraw_InsufficientFunds(t *testing.T) {
	l := NewLedgerStore()
	l.Deposit(50)

	err := l.Withdraw(100)
	assert.Error(t, err, "Withdraw should fail if funds are insufficient")
	assert.Equal(t, 50, l.GetBalance().Amount, "Balance should remain unchanged on failed withdrawal")
	assert.Len(t, l.GetTransactions().Transactions, 1, "Failed withdrawal should not be recorded")
}

func TestConcurrentAccess(t *testing.T) {
	l := NewLedgerStore()
	done := make(chan bool)

	// Run concurrent deposits
	for i := 0; i < 10; i++ {
		go func() {
			l.Deposit(10)
			done <- true
		}()
	}

	// Run concurrent withdrawals
	for i := 0; i < 5; i++ {
		go func() {
			l.Withdraw(10)
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 15; i++ {
		<-done
	}

	assert.Equal(t, 50, l.GetBalance().Amount, "Final balance should be consistent with concurrent operations")
	assert.Len(t, l.GetTransactions().Transactions, 15, "All transactions should be recorded correctly")
}
