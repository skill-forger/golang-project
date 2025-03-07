package profile

import (
	ct "golang-project-layout/internal/contract"
	repo "golang-project-layout/internal/repository"
	svc "golang-project-layout/internal/service"
)

// service represents the implementation of service.Profile
type service struct {
	userRepo repo.User
}

// NewService returns a new implementation of service.Profile
func NewService(userRepo repo.User) svc.Profile {
	return &service{
		userRepo: userRepo,
	}
}

// GetByID executes the profile detail retrieval logic
func (s *service) GetByID(id int) (*ct.ProfileResponse, error) {
	user, err := s.userRepo.Read(id)
	if err != nil {
		return nil, err
	}

	return prepareProfileResponse(user), nil
}
