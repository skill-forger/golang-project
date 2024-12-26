package database

import (
	"errors"

	"gorm.io/gorm"
)

func RunInTx(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := fn(tx)
	if err == nil {
		commitResult := tx.Commit()
		return commitResult.Error
	}

	rollbackResult := tx.Rollback()
	if rollbackResult.Error != nil {
		return errors.Join(err, rollbackResult.Error)
	}

	return err
}
