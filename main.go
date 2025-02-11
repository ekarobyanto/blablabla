package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/username/mentoring_study_case/auth"
	"github.com/username/mentoring_study_case/db"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	con := db.InitDB()
	log.Println("Database connection established")
	app := fiber.New()
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	app.Use(cors.New())
	app.Group("/api")
	app.Group("/v1")

	authHandler := auth.InitializeAuthHandler(con)
	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)
}
