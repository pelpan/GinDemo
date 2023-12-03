package model

type RequestsBody struct {
	Name    string   `json:"name"`
	BankNos []string `json:"bankNos"`
}