package data

import (
	"errors"
	"technopartner/features/outcomes"

	"gorm.io/gorm"
)

type OutcomeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) outcomes.OutcomeDataInterface {
	return &OutcomeQuery{
		db: db,
	}
}

func (repo *OutcomeQuery) Insert(UserID uint, input outcomes.OutcomeCore) error {
	var outcomeModel = OutcomeCoreToModel(input)

	tx := repo.db.Where("user_id = ?", UserID).Create(&outcomeModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *OutcomeQuery) GetAllOutcomes(UserID uint) ([]outcomes.OutcomeCore, error) {
	var outcomes []outcomes.OutcomeCore
	if err := repo.db.Where("user_id = ?", UserID).Find(&outcomes).Error; err != nil {
		return nil, err
	}
	return outcomes, nil
}

func (repo *OutcomeQuery) DeleteOutcome(UserID, OutcomeID uint) error {
	tx := repo.db.Where("id = ? AND tenant_id = ?", OutcomeID, UserID).Delete(&Outcome{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *OutcomeQuery) UpdateOutcome(UserID, OutcomeID uint, input outcomes.OutcomeCore) error {
	var outcome Outcome
	tx := repo.db.Where("id = ? AND tenant_id = ?", OutcomeID, UserID).First(&outcome)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	updatedOutcome := OutcomeCoreToModel(input)

	tx = repo.db.Model(&outcome).Updates(updatedOutcome)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}
