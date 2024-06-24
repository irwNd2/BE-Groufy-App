package main

import (
	"be-groufy-app/handlers"
	"be-groufy-app/storage"
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

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Database connection error")
	}
	err = storage.Migrate(db)

	if err != nil {
		log.Fatal("Database migrating error")
	}

	app := fiber.New()
	handlers.SetupInfoRoutes(app, db)
	handlers.SetupUserRoutes(app, db)

	app.Listen(":3000")
}
