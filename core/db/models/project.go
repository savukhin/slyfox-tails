package models

type Project struct {
	ID        uint64 `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Title     string `gorm:"column:title;type:text;not null" json:"title"`
	CreatorID uint64 `gorm:"column:creator_id;type:bigint;not null" json:"creator_id"`
	Jobs      []Job  `gorm:"foreignKey:ProjectID" json:"jobs"`

	TimingAt
}
