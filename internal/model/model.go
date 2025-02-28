package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int `gorm:"TYPE:BIGINT(11);UNSIGNED;AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}
