package incomes

import "time"

type IncomeCore struct {
	ID        uint
	Name      string
	Jumlah    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type IncomeDataInterface interface {
	Insert(UserID uint, input IncomeCore) error
	GetAllIncomes(UserID uint) ([]IncomeCore, error)
	DeleteIncome(UserID, IncomeID uint) error
	UpdateIncome(UserID, IncomeID uint, input IncomeCore) error
}

type IncomeServiceInterface interface {
	CreateIncome(UserID uint, input IncomeCore) error
	GetAllIncomes(UserID uint) ([]IncomeCore, error)
	DeleteIncome(UserID, IncomeID uint) error
	UpdateIncome(UserID, IncomeID uint, input IncomeCore) error
}
