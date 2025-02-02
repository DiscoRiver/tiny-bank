package handlers

import (
	"encoding/json"
	"github.com/discoriver/tiny-bank/ledgerstore"
	"net/http"
	"strconv"
)

// Global ledger instance
var (
	l = ledgerstore.NewLedgerStore()
)

// HandleDeposit handles deposit requests
func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	l.Deposit(amount)

	w.WriteHeader(http.StatusOK)
}

// HandleWithdraw handles withdrawal requests
func HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	err = l.Withdraw(amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleBalance returns the current balance
func HandleBalance(w http.ResponseWriter, r *http.Request) {
	balance := l.GetBalance()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

// HandleTransactions returns transaction history
func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := l.GetTransactions()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
