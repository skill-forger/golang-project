package contract

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaim struct {
	jwt.StandardClaims
	UserID    int    `json:"user_id,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
}

type ContextUser struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignInResponse struct {
	UserID       int    `json:"user_id,omitempty"`
	Token        string `json:"token,omitempty"`
	Type         string `json:"type,omitempty"`
	ExpiredAfter int    `json:"expired_at,omitempty"`
}
