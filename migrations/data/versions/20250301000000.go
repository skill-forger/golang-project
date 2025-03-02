package versions

import (
	"gorm.io/gorm"

	"golang-project-layout/internal/model"
	"golang-project-layout/util/hashing"
)

func Migrate20250301000000(tx *gorm.DB) error {
	demoPassword, err := hashing.NewBcrypt().Generate([]byte("demouser@123"))
	if err != nil {
		return err
	}

	type User struct {
		model.BaseModel
		FirstName    string
		LastName     string
		Email        string
		Password     string
		DisplayName  string
		ProfileImage string
		Biography    string
	}

	data := &User{
		FirstName:    "demo",
		LastName:     "user",
		Email:        "user@demo.com",
		Password:     string(demoPassword),
		DisplayName:  "demo_user",
		ProfileImage: "demo user profile image",
		Biography:    "this is the demo user biography",
	}

	return tx.Model(&User{}).Create(data).Error
}
