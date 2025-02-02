package main

import (
	"github.com/discoriver/tiny-bank/handlers"
	"log"
	"net/http"
)

func main() {
	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/deposit", handlers.HandleDeposit)
	mux.HandleFunc("/withdraw", handlers.HandleWithdraw)
	mux.HandleFunc("/balance", handlers.HandleBalance)
	mux.HandleFunc("/transactions", handlers.HandleTransactions)

	// Start
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
