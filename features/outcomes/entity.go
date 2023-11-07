package outcomes

import "time"

type OutcomeCore struct {
	ID        uint
	Name      string
	Jumlah    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type OutcomeDataInterface interface {
	Insert(UserID uint, input OutcomeCore) error
	GetAllOutcomes(UserID uint) ([]OutcomeCore, error)
	DeleteOutcome(UserID, OutcomeID uint) error
	UpdateOutcome(UserID, OutcomeID uint, input OutcomeCore) error
}

type OutcomeServiceInterface interface {
	CreateOutcome(UserID uint, input OutcomeCore) error
	GetAllOutcomes(UserID uint) ([]OutcomeCore, error)
	DeleteOutcome(UserID, OutcomeID uint) error
	UpdateOutcome(UserID, OutcomeID uint, input OutcomeCore) error
}
