package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int `json:"toAccount`
	Amout     int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreateAt  time.Time `json:"createAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(10000000)),
		Balance:   0,
		CreateAt:  time.Now().UTC(),
	}
}
