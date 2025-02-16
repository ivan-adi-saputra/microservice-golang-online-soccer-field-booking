package seeders

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	user := models.User{
		UUID:        uuid.New(),
		Name:        "Administrator",
		Username:    "admin",
		Email:       "admin@gmail.com",
		Password:    string(password),
		PhoneNumber: "081234567",
		RoleID:      constants.Admin,
	}

	if err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error; err != nil {
		logrus.Errorf("failed to seed user: %v", err)
		panic(err)
	}

	logrus.Infof("user %s successfully seeded", user.Username)
}
