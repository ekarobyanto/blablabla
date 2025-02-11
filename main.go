package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"mentoring_study_case/db"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	_ = db.InitDB()

	log.Println("Database connection established")

	app := fiber.New()
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
