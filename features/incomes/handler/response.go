package handler

import "technopartner/features/incomes"

type IncomeResponseAll struct {
	Name   string `json:"name"`
	Jumlah uint   `json:"jumlah"`
}

func IncomeCoreToResponseAll(input incomes.IncomeCore) IncomeResponseAll {
	var incomeResp = IncomeResponseAll{
		Name:   input.Name,
		Jumlah: input.Jumlah,
	}
	return incomeResp
}
