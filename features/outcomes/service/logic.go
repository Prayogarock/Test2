package service

import (
	"technopartner/features/outcomes"
)

type OutcomeService struct {
	outcomeData outcomes.OutcomeDataInterface
}

func New(repo outcomes.OutcomeDataInterface) outcomes.OutcomeServiceInterface {
	return &OutcomeService{
		outcomeData: repo,
	}
}

func (service *OutcomeService) CreateOutcome(UserID uint, input outcomes.OutcomeCore) error {
	return service.outcomeData.Insert(UserID, input)
}

func (service *OutcomeService) GetAllOutcomes(UserID uint) ([]outcomes.OutcomeCore, error) {
	return service.outcomeData.GetAllOutcomes(UserID)
}

func (service *OutcomeService) DeleteOutcome(UserID, OutcomeID uint) error {
	return service.outcomeData.DeleteOutcome(UserID, OutcomeID)
}

func (service *OutcomeService) UpdateOutcome(UserID, OutcomeID uint, input outcomes.OutcomeCore) error {
	return service.outcomeData.UpdateOutcome(UserID, OutcomeID, input)
}
