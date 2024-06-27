package main

import (
	apps "be-groufy-app/app"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &apps.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := apps.NewConnection(config)
	if err != nil {
		log.Fatal("Database connection error")
	}
	err = apps.Migrate(db)

	if err != nil {
		log.Fatal("Database migrating error")
	}
	app := fiber.New()

	apps.SetupUserRoutes(app, db)
	apps.SetupInfoRoutes(app, db)

	app.Listen(":3000")
}
