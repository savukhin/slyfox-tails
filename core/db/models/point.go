package models

type Point struct {
	ID        uint64   `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Title     string   `gorm:"column:title;type:text;not null" json:"title"`
	CreatorID uint64   `gorm:"column:creator_id;type:bigint;not null" json:"creator_id"`
	Stages    []*Stage `gorm:"many2many:point_stage" json:"stages"`

	TimingAt
}
