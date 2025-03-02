package data

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"golang-project-layout/migrations/data/versions"
)

func Migrate(db *gorm.DB) error {
	option := gormigrate.DefaultOptions
	option.TableName = "data_migrations"

	m := gormigrate.New(db, option, []*gormigrate.Migration{
		{
			ID:      "20250301000000",
			Migrate: versions.Migrate20250301000000,
		},
	})

	return m.Migrate()
}
