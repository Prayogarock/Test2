package data

import (
	_income "technopartner/features/incomes/data"
	_outcome "technopartner/features/outcomes/data"
	"technopartner/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"name;not null"`
	Email       string `gorm:"email;not null"`
	Password    string `gorm:"password;not null"`
	PhoneNumber string `gorm:"phone_number;not null"`
	Address     string `gorm:"address;not null"`
	Saldo       uint   `gorm:"saldo;not null"`
	Incomes     []_income.Income
	Outcomes    []_outcome.Outcome
}

func ModelToUserCore(input User) users.UserCore {
	return users.UserCore{
		ID:          input.ID,
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Saldo:       input.Saldo,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
}

func UserCoreToModel(input users.UserCore) User {
	var userModel = User{
		Model:       gorm.Model{},
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Saldo:       input.Saldo,
	}
	return userModel
}
