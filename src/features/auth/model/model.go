package auth_model

import "time"

type User struct {
	Id              int        `gorm:"column:id;" json:"id"`
	Username        string     `gorm:"column:username;" json:"username"`
	Password        string     `gorm:"column:password;" json:"-"`
	Email           string     `gorm:"column:email;" json:"email"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at" json:"email_verified_at"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
