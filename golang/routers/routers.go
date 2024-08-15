// routers/routers.go
package routers

import (
	"brandAPI/internal/handlers"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the application routes
func SetupRoutes(app *fiber.App, db *sql.DB) {
	// USER routes
	app.Post("/users", handlers.CreateUserHandler(db))
	app.Get("/users", handlers.GetAllUserHandler(db))
	app.Patch("/users/:id", handlers.UpdateUserHandler(db))
}
