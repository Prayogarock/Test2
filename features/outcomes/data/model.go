package data

import (
	"technopartner/features/outcomes"

	"gorm.io/gorm"
)

type Outcome struct {
	gorm.Model
	Name   string `gorm:"name;not null"`
	Jumlah uint   `gorm:"jumlah;not null"`
	UserID uint
}

func OutcomeCoreToModel(input outcomes.OutcomeCore) Outcome {
	var outcomeModel = Outcome{
		Model:  gorm.Model{},
		Name:   input.Name,
		Jumlah: input.Jumlah,
	}
	return outcomeModel
}
