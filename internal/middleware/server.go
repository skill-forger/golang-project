package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Timeout() echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 60 * time.Second})
}

func Recover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(e echo.Context, err error, stack []byte) error {
			fmt.Println("panic error:", err)
			return nil
		},
	})
}

func ErrorHandler(err error, e echo.Context) {
	if e.Response().Committed {
		return
	}

	httpCode := http.StatusInternalServerError
	message := http.StatusText(http.StatusInternalServerError)
	cid := e.Response().Header().Get(echo.HeaderXCorrelationID)

	httpErr, ok := err.(*echo.HTTPError)
	if ok {
		httpCode = httpErr.Code
		message = fmt.Sprintf("%s", httpErr.Message)
	}

	_ = e.JSON(httpCode, map[string]any{"cid": cid, "message": message})
}
