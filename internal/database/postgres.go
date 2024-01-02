package database

import (
	"fmt"
	"log"
	"os"

	"github.com/tfkhdyt/yukitanya-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Question{},
		&model.Answer{},
		&model.Subject{},
		&model.OldSlug{},
		&model.Membership{},
		&model.Notification{},
		&model.QuestionImage{},
		&model.Rating{},
	); err != nil {
		log.Fatalln("Error:", err)
	}

	return db
}
