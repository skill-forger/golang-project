package database

import (
	"errors"

	"gorm.io/gorm"
)

// RunInTx wrap the logic function inside a gorm transaction
func RunInTx(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := fn(tx)
	if err == nil {
		return tx.Commit().Error
	}

	rollbackResult := tx.Rollback()
	if rollbackResult.Error != nil {
		return errors.Join(err, rollbackResult.Error)
	}

	return err
}
