package handler

import "technopartner/features/incomes"

type IncomeRequest struct {
	Name   string `json:"name" form:"name"`
	Jumlah uint   `json:"jumlah" form:"jumlah"`
}

func IncomeRequestToCore(input IncomeRequest) incomes.IncomeCore {
	var incomeCore = incomes.IncomeCore{
		Name:   input.Name,
		Jumlah: input.Jumlah,
	}
	return incomeCore
}
