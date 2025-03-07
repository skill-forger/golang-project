package authentication

import (
	"github.com/spf13/viper"

	ct "golang-project-layout/internal/contract"
	m "golang-project-layout/internal/model"
	"golang-project-layout/static"
)

// prepareSignInResponse transforms the data and returns the Sign In Response
func prepareSignInResponse(o *m.User, token string) *ct.SignInResponse {
	return &ct.SignInResponse{
		UserID:       o.ID,
		Token:        token,
		Type:         viper.GetString(static.EnvAuthType),
		ExpiredAfter: viper.GetInt(static.EnvAuthLifeTime),
	}
}
