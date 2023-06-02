package models

type Job struct {
	ID        uint64  `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Title     string  `gorm:"column:title;type:text;not null" json:"title"`
	CreatorID uint64  `gorm:"column:creator_id;type:bigint;not null" json:"creator_id"`
	ProjectID uint64  `gorm:"column:project_id;type:bigint;not null" json:"project_id"`
	Stages    []Stage `gorm:"foreignKey:JobID" json:"stages"`

	TimingAt
}
