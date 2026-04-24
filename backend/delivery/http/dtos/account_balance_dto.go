package dtos

type AccountBalanceDTO struct {
	AccountID string  `json:"accountID"` //PublicID
	Balance   float64 `json:"balance"`
}
