package handler

import "technopartner/features/outcomes"

type OutcomeRequest struct {
	Name   string `json:"name" form:"name"`
	Jumlah uint   `json:"jumlah" form:"jumlah"`
}

func OutcomeRequestToCore(input OutcomeRequest) outcomes.OutcomeCore {
	var outcomeCore = outcomes.OutcomeCore{
		Name:   input.Name,
		Jumlah: input.Jumlah,
	}
	return outcomeCore
}
