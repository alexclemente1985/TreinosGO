package dtos

type UserDTO struct {
	ID        string `json:"id"` //PublicID
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}
