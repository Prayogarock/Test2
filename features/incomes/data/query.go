package data

import (
	"errors"
	"technopartner/features/incomes"

	"gorm.io/gorm"
)

type IncomeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) incomes.IncomeDataInterface {
	return &IncomeQuery{
		db: db,
	}
}

func (repo *IncomeQuery) Insert(UserID uint, input incomes.IncomeCore) error {
	var incomeModel = IncomeCoreToModel(input)

	tx := repo.db.Where("user_id = ?", UserID).Create(&incomeModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *IncomeQuery) GetAllIncomes(UserID uint) ([]incomes.IncomeCore, error) {
	var incomes []incomes.IncomeCore
	if err := repo.db.Where("user_id = ?", UserID).Find(&incomes).Error; err != nil {
		return nil, err
	}
	return incomes, nil
}

func (repo *IncomeQuery) DeleteIncome(UserID, IncomeID uint) error {
	var income Income
	tx := repo.db.Where("id = ? AND tenant_id = ?", IncomeID, UserID).First(&income)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	tx = repo.db.Where("id = ? AND tenant_id = ?", IncomeID, UserID).Delete(&income)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *IncomeQuery) UpdateIncome(UserID, IncomeID uint, input incomes.IncomeCore) error {
	var income Income
	tx := repo.db.Where("id = ? AND tenant_id = ?", IncomeID, UserID).First(&income)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	updatedIncome := IncomeCoreToModel(input)

	tx = repo.db.Model(&income).Updates(updatedIncome)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}
