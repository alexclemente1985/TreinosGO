package entities

import "time"

type User struct {
	ID        int64  `gorm:"primaryKey"`
	PublicID  string `gorm:"type:char(26);uniqueIndex"`
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
}
