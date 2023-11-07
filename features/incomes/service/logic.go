package service

import (
	"technopartner/features/incomes"
)

type IncomeService struct {
	incomeData incomes.IncomeDataInterface
}

func New(repo incomes.IncomeDataInterface) incomes.IncomeServiceInterface {
	return &IncomeService{
		incomeData: repo,
	}
}

func (service *IncomeService) CreateIncome(UserID uint, input incomes.IncomeCore) error {
	return service.incomeData.Insert(UserID, input)
}

func (service *IncomeService) GetAllIncomes(UserID uint) ([]incomes.IncomeCore, error) {
	return service.incomeData.GetAllIncomes(UserID)
}

func (service *IncomeService) DeleteIncome(UserID, IncomeID uint) error {
	return service.incomeData.DeleteIncome(UserID, IncomeID)
}

func (service *IncomeService) UpdateIncome(UserID, IncomeID uint, input incomes.IncomeCore) error {
	return service.incomeData.UpdateIncome(UserID, IncomeID, input)
}
