package registry

import (
	"gorm.io/gorm"

	"golang-project-layout/server"

	"golang-project-layout/internal/handler"
	"golang-project-layout/internal/registry/authentication"
	"golang-project-layout/internal/registry/profile"
)

func initResourceHandlers(db *gorm.DB) []handler.ResourceHandler {
	return []handler.ResourceHandler{
		authentication.NewRegistry("/auth", db),
		profile.NewRegistry("/profile", db),
	}
}

func NewHandlerRegistries(db *gorm.DB) []server.HandlerRegistry {
	var registries []server.HandlerRegistry

	for _, hdl := range initResourceHandlers(db) {
		registries = append(registries, hdl.RegisterRoutes())
	}

	return registries
}
