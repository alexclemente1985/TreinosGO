package entities

type Account struct {
	ID         int64  `gorm:"primaryKey"`                //interno
	PublicID   string `gorm:"type:char(26);uniqueIndex"` //externo (ULID)
	CustomerID int64  `gorm:"index;not null"`
	Type       string // checking | savings
	Balance    float64
}
