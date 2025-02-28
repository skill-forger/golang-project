package service

import (
	ct "golang-project-layout/internal/contract"
)

type Authentication interface {
	SignIn(*ct.SignInRequest) (*ct.SignInResponse, error)
}

type Profile interface {
	GetDetail()
}
