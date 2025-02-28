package health

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"golang-project-layout/database"
	"golang-project-layout/internal/contract"
	"golang-project-layout/server"
)

func NewRegistry(dbConnection database.Connection) server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: "/health",
		Register: func(group *echo.Group) {
			group.GET("", func(e echo.Context) error {
				return e.JSON(http.StatusOK, prepareHealthCheckResponse(dbConnection))
			})
		},
	}
}

func prepareHealthCheckResponse(dbConnection database.Connection) []*contract.HealthCheckResponse {
	dbMessage := "ok"
	err := dbConnection.Ping()
	if err != nil {
		dbMessage = fmt.Sprintf("error: %s", err)
	}

	return []*contract.HealthCheckResponse{
		{
			Resource: "server",
			Status:   "ok",
		},
		{
			Resource: "database",
			Status:   dbMessage,
		},
	}
}
