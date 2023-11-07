package data

import (
	"technopartner/features/incomes"

	"gorm.io/gorm"
)

type Income struct {
	gorm.Model
	Name   string `gorm:"name;not null"`
	Jumlah uint   `gorm:"jumlah;not null"`
	UserID uint
}

func IncomeCoreToModel(input incomes.IncomeCore) Income {
	var incomeModel = Income{
		Model:  gorm.Model{},
		Name:   input.Name,
		Jumlah: input.Jumlah,
	}
	return incomeModel
}
