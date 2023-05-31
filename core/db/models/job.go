package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID        uint64  `gorm:"column:id;type:serial;primaryKey" json:"id"`
	Title     string  `gorm:"column:title;type:text" json:"title"`
	CreatorID uint64  `gorm:"column:creator_id;type:serial" json:"creator_id"`
	ProjectID uint64  `gorm:"column:project_id;type:serial" json:"project_id"`
	Stages    []Stage `gorm:"foreignKey:JobID" json:"stages"`

	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_attt;type:timestamp with time zone" json:"deleted_at"`
}
