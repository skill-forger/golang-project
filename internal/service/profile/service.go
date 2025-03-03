package profile

import (
	ct "golang-project-layout/internal/contract"
	repo "golang-project-layout/internal/repository"
	svc "golang-project-layout/internal/service"
)

type service struct {
	userRepo repo.User
}

func NewService(userRepo repo.User) svc.Profile {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) GetByID(id int) (*ct.ProfileResponse, error) {
	user, err := s.userRepo.Read(id)
	if err != nil {
		return nil, err
	}

	return prepareProfileResponse(user), nil
}
