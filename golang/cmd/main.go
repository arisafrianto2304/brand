package main

import (
	"brandAPI/configs"
	_ "brandAPI/docs"
	routers "brandAPI/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()
	log.Println("Connecting to the database...")
	db := configs.ConnectDB()
	defer db.Close()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Set up routes
	routers.SetupRoutes(app, db)

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "http://example.com/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
