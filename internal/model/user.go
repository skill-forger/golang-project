package model

type User struct {
	BaseModel
	Email    string
	Password string
	Name     string
}
