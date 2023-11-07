package handler

import "technopartner/features/users"

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Saldo       uint   `json:"saldo" form:"saldo"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRequestToCore(input UserRequest) users.UserCore {
	var userCore = users.UserCore{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Saldo:       input.Saldo,
	}
	return userCore
}
