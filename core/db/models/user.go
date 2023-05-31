package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int       `gorm:"column:id;type:serial;primaryKey" json:"id"`
	Username string    `gorm:"column:username;type:varchar(20)" json:"username"`
	Email    string    `gorm:"column:email;type:varchar(20)" json:"email"`
	Points   []Point   `gorm:"foreignKey:CreatorID" json:"points"`
	Projects []Project `gorm:"foreignKey:CreatorID" json:"projects"`

	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_attt;type:timestamp with time zone" json:"deleted_at"`
}
