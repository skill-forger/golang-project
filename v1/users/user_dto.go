package users

type CreatedUserDTO struct {
	Username string `json:"username" validate:"required,min=1,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1"`
}

type UserDetailDTO struct {
	Username string `json:"username" validate:"required,min=1,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1"`
}
