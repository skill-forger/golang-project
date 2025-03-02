package schema

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"

	"golang-project-layout/migrations/schema/versions"
)

func Migrate(db *gorm.DB) error {
	option := gormigrate.DefaultOptions
	option.TableName = "schema_migrations"

	m := gormigrate.New(db, option, []*gormigrate.Migration{
		{
			ID:       "20250301000000",
			Migrate:  versions.Migrate20250301000000,
			Rollback: versions.Rollback20250301000000,
		},
	})

	return m.Migrate()
}
