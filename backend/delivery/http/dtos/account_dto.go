package dtos

type AccountDTO struct {
	ID         string  `json:"id"`         //PublicID
	CustomerID string  `json:"customerID"` //PublicID do Customer
	Type       string  `json:"type"`
	Balance    float64 `json:"balance"`
}
