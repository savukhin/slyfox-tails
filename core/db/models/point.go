package models

type Point struct {
	ID           uint64   `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	PasswordHash string   `gorm:"column:password_hash;type:varchar(72);not null" json:"password_hash"`
	Login        string   `gorm:"column:login;type:text;not null" json:"login"`
	Title        string   `gorm:"column:title;type:text;not null" json:"title"`
	CreatorID    uint64   `gorm:"column:creator_id;type:bigint;not null" json:"creator_id"`
	Stages       []*Stage `gorm:"many2many:point_stage" json:"stages"`

	TimingAt
}
