package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"
	"golang-project/util/hashing"
)

// service represents the implementation of service.Authentication
type service struct {
	userRepo repo.User
	hash     hashing.Algorithm
}

// NewService returns a new implementation of service.Authentication
func NewService(userRepo repo.User, hash hashing.Algorithm) svc.Authentication {
	return &service{
		userRepo: userRepo,
		hash:     hash,
	}
}

// SignIn executes the user authentication logic
func (s *service) SignIn(r *ct.SignInRequest) (*ct.SignInResponse, error) {
	user, err := s.userRepo.ReadByEmail(r.Email)
	if err != nil {
		return nil, err
	}

	err = s.hash.Compare([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return prepareSignInResponse(user, token), nil
}

// generateToken returns the JWT token based on the information from model.User
func (s *service) generateToken(user *model.User) (string, error) {
	secret := []byte(viper.GetString(static.EnvAuthSecret))
	customClaim := &ct.CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  []string{viper.GetString(static.EnvAuthAudience)},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(viper.GetInt64(static.EnvAuthLifeTime)) * time.Second)),
			ID:        uuid.NewString(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    viper.GetString(static.EnvAuthIssuer),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   viper.GetString(static.EnvAuthSubject),
		},
		UserID:    user.ID,
		UserEmail: user.Email,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim).SignedString(secret)
}
