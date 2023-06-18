package models

import "time"

type Stage struct {
	ID        uint64    `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Title     string    `gorm:"column:title;type:text;not null" json:"title"`
	CreatorID uint64    `gorm:"column:creator_id;type:bigint;not null" json:"creator_id"`
	JobID     uint64    `gorm:"column:job_id;type:bigint" json:"job_id"`
	Points    []*Point  `gorm:"many2many:point_stage;" json:"points"`
	StartedAt time.Time `gorm:"column:started_at;type:timestamp with time zone;" json:"started_at"`

	TimingAt
}
