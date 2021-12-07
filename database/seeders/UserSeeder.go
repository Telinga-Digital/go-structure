package seeders

import (
	"github.com/Telinga-Digital/go-structure/app/models"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
)

var (
	count int
	users []models.User
)

func UserSeeder() {
	count = 10

	for i := 1; i < count; i++ {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

		user := models.User{
			Email:    faker.Email(),
			Password: string(hashedPassword),
			ApiKey:   faker.UUIDDigit(),
		}

		users = append(users, user)
	}

	models.DB.Create(&users)
}
