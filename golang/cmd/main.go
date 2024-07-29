package main

import (
	"brandAPI/configs"
	_ "brandAPI/docs"
	"brandAPI/internal/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()
	log.Println("Connecting to the database...")
	db := configs.ConnectDB()
	defer db.Close()

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	// Register POST /users route
	app.Post("/users", handlers.CreateUserHandler(db))
	app.Get("/users", handlers.GetAllUserHandler(db))

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
