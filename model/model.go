package model

type Balance struct {
	Amount int `json:"amount"`
}

type Transaction struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}
