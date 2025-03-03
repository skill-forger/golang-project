package model

type User struct {
	BaseModel
	FirstName    string
	LastName     string
	Email        string
	Password     string
	DisplayName  string
	ProfileImage string
	Biography    string
}
