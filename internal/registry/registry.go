package registry

import (
	"gorm.io/gorm"

	"golang-project-layout/database"
	"golang-project-layout/internal/handler"
	"golang-project-layout/internal/registry/authentication"
	"golang-project-layout/internal/registry/health"
	"golang-project-layout/internal/registry/profile"
	"golang-project-layout/server"
)

func initResourceHandlers(dbInstance *gorm.DB) []handler.ResourceHandler {
	return []handler.ResourceHandler{
		authentication.NewRegistry("/auth", dbInstance),
		profile.NewRegistry("/profile", dbInstance),
	}
}

func NewHandlerRegistries(dbConnection database.Connection) ([]server.HandlerRegistry, error) {
	dbInstance, err := dbConnection.Instance()
	if err != nil {
		return nil, err
	}

	registries := []server.HandlerRegistry{health.NewRegistry(dbConnection)}

	for _, hdl := range initResourceHandlers(dbInstance) {
		registries = append(registries, hdl.RegisterRoutes())
	}

	return registries, nil
}
