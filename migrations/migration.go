package migrations

import (
	"golang-project-layout/config"
	"golang-project-layout/internal/model"
)

func Migrate(app *config.AppConfig) {
	db, err := app.DB.Instance()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
