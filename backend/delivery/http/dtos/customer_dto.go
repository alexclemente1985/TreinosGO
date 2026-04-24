package dtos

type CustomerDTO struct {
	ID          string `json:"id"` //PublicID
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	MobilePhone string `json:"mobilePhone"`
	CreatedAt   string `json:"createdAt"`
}
