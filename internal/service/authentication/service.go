package authentication

import (
	repo "golang-project-layout/internal/repository"
	svc "golang-project-layout/internal/service"
)

type service struct {
	userRepo repo.User
}

func NewService(userRepo repo.User) svc.Authentication {
	return &service{
		userRepo: userRepo,
	}
}

func (p *service) SignIn() {
	//TODO implement me
	panic("implement me")
}
