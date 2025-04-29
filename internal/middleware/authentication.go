package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	ct "golang-project/internal/contract"
	"golang-project/static"
)

// Authentication provides the middleware for any API requires user authentication
func Authentication() echo.MiddlewareFunc {
	return echoJwt.WithConfig(echoJwt.Config{
		Skipper:       func(c echo.Context) bool { return false },
		SigningKey:    []byte(viper.GetString(static.EnvAuthSecret)),
		SigningMethod: echoJwt.AlgorithmHS256,
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString(static.EnvAuthSecret)), nil
			}

			token, err := jwt.ParseWithClaims(auth, &ct.CustomClaim{}, keyFunc)
			if err != nil {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			claim, ok := token.Claims.(*ct.CustomClaim)
			if !ok || !token.Valid {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "parse jwt custom claim failed")
			}

			return &ct.ContextUser{ID: claim.UserID, Email: claim.UserEmail}, nil
		},
	})
}
