package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(db *gorm.DB) {
	if err := recover(); err != nil {
		db.Rollback()
	} else {
		db.Commit()
	}
}
