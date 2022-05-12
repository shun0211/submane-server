package domain

type Users []User

type User struct {
	ID int
	Name string `json:"name" validate:"required,name"`
}
