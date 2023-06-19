package models

type User struct {
	ID            uint64    `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Username      string    `gorm:"column:username;type:varchar(20);not null;unique" json:"username"`
	PasswordHash  string    `gorm:"column:password_hash;type:varchar(72);not null;unique" json:"password_hash"`
	Email         string    `gorm:"column:email;type:varchar(20);not null;unique" json:"email"`
	EmailVerified bool      `gorm:"column:email_verified;type:bool;not null;default: false" json:"email_verified"`
	Points        []Point   `gorm:"foreignKey:CreatorID" json:"points"`
	Projects      []Project `gorm:"foreignKey:CreatorID" json:"projects"`
	Jobs          []Job     `gorm:"foreignKey:CreatorID" json:"jobs"`
	Stages        []Stage   `gorm:"foreignKey:CreatorID" json:"stages"`

	TimingAt
}
