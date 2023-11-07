package service

import (
	"technopartner/app/middlewares"
	"technopartner/features/users"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
		validate: validator.New(),
	}
}

func (service *UserService) Login(email string, password string) (dataLogin users.UserCore, token string, err error) {
	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateTokenUser(dataLogin.ID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}

func (service *UserService) Create(input users.UserCore) (dataCreate users.UserCore, token string, err error) {
	dataCreate, err = service.userData.Insert(input)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateTokenUser(dataCreate.ID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataCreate, token, nil
}
