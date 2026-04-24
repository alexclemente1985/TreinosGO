package entities

import "time"

type Customer struct {
	ID          int64  `gorm:"primaryKey"`                //interno
	PublicID    string `gorm:"type:char(26);uniqueIndex"` //externo (ULID)
	Name        string
	Email       string
	Phone       string
	MobilePhone string
	CreatedAt   time.Time
}
