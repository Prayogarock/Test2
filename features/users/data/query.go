package data

import (
	"errors"
	"technopartner/features/users"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db:        db,
		dataLogin: users.UserCore{},
	}
}

func (repo *UserQuery) Login(email string, password string) (dataLogin users.UserCore, err error) {

	var data User
	tx := repo.db.Where("email = ?", email).Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}
	dataLogin = ModelToUserCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}

func (repo *UserQuery) Insert(input users.UserCore) (dataCreate users.UserCore, err error) {
	var userModel = UserCoreToModel(input)

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	return users.UserCore{}, tx.Error
}
