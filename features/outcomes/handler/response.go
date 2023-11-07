package handler

import "technopartner/features/outcomes"

type OutcomeResponseAll struct {
	Name   string `json:"name"`
	Jumlah uint   `json:"jumlah"`
}

func OutcomeCoreToResponseAll(input outcomes.OutcomeCore) OutcomeResponseAll {
	var outcomeResp = OutcomeResponseAll{
		Name:   input.Name,
		Jumlah: input.Jumlah,
	}
	return outcomeResp
}
