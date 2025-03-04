package authentication

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	ct "golang-project-layout/internal/contract"
	"golang-project-layout/internal/model"
	repo "golang-project-layout/internal/repository"
	svc "golang-project-layout/internal/service"
	"golang-project-layout/static"
	"golang-project-layout/util/hashing"
)

type service struct {
	userRepo repo.User
	hash     hashing.Algorithm
}

func NewService(userRepo repo.User, hash hashing.Algorithm) svc.Authentication {
	return &service{
		userRepo: userRepo,
		hash:     hash,
	}
}

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

func (s *service) generateToken(user *model.User) (string, error) {
	secret := []byte(viper.GetString(static.EnvAuthSecret))
	customClaim := &ct.CustomClaim{
		StandardClaims: jwt.StandardClaims{
			Audience:  viper.GetString(static.EnvAuthAudience),
			ExpiresAt: time.Now().Unix() + viper.GetInt64(static.EnvAuthLifeTime),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    viper.GetString(static.EnvAuthIssuer),
			NotBefore: time.Now().Unix(),
			Subject:   viper.GetString(static.EnvAuthSubject),
		},
		UserID:    user.ID,
		UserEmail: user.Email,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim).SignedString(secret)
}
