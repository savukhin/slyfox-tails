package models

import (
	"time"

	"gorm.io/gorm"
)

type TimingAt struct {
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
}
