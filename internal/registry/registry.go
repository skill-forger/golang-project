package registry

import (
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	_ "golang-project-layout/docs/swagger"
	"golang-project-layout/internal/handler"
	"golang-project-layout/internal/registry/authentication"
	"golang-project-layout/internal/registry/health"
	"golang-project-layout/internal/registry/profile"
	"golang-project-layout/server"
)

func NewHandlerRegistries(db *gorm.DB) ([]server.HandlerRegistry, error) {
	registries := []server.HandlerRegistry{
		initSwaggerRegistry(),
		initHealthCheckHandler(db).RegisterRoutes(),
	}

	for _, hdl := range initResourceHandlers(db) {
		registries = append(registries, hdl.RegisterRoutes())
	}

	return registries, nil
}

func initSwaggerRegistry() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: "/swagger",
		Register: func(group *echo.Group) {
			group.GET("/*", swagger.WrapHandler)
		},
	}
}

func initHealthCheckHandler(db *gorm.DB) handler.ResourceHandler {
	return health.NewRegistry("/health", db)
}

func initResourceHandlers(db *gorm.DB) []handler.ResourceHandler {
	return []handler.ResourceHandler{
		authentication.NewRegistry("/auth", db),
		profile.NewRegistry("/profile", db),
	}
}
