package models

type User struct {
	ID       int       `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Username string    `gorm:"column:username;type:varchar(20);not null" json:"username"`
	Email    string    `gorm:"column:email;type:varchar(20);not null" json:"email"`
	Points   []Point   `gorm:"foreignKey:CreatorID" json:"points"`
	Projects []Project `gorm:"foreignKey:CreatorID" json:"projects"`

	TimingAt
}
