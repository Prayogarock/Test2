package users

import (
	_incomeCore "technopartner/features/incomes"
	_outcomeCore "technopartner/features/outcomes"
	"time"
)

type UserCore struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	Saldo       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Incomes     []_incomeCore.IncomeCore
	Outcomes    []_outcomeCore.OutcomeCore
}

type UserDataInterface interface {
	Login(email, password string) (UserCore, error)
	Insert(input UserCore) (UserCore, error)
}

type UserServiceInterface interface {
	Login(email, password string) (UserCore, string, error)
	Create(input UserCore) (UserCore, string, error)
}
