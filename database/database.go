package database

import (
	"os"
	util "pelatihan-be/helpers/utils"
	"pelatihan-be/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- util.GodotEnv("DB_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DB_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&model.UserLoginEntityModel{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
