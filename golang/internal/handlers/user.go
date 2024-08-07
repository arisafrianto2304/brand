package handlers

import (
	"brandAPI/internal/models"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserHandler creates a user with the provided details
// @Summary Create a new user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User required "User Data"
// @Success 201 {object} models.User
// @Failure 400 {string} string "invalid input, object invalid"
// @Failure 500 {string} string "internal server error"
// @Router /users [post]
func CreateUserHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)

		if err := c.BodyParser(user); err != nil {
			log.Println("Failed to parse user:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		// Hashing the password before storing it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Failed to hash password:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to secure user data",
			})
		}

		// Set current time as created_at
		currentTime := time.Now()

		// Use the correct table name 'user' instead of 'users'
		_, err = db.Exec(`INSERT INTO "user" (username, password, created_at) VALUES ($1, $2, $3)`, user.Username, string(hashedPassword), currentTime)
		if err != nil {
			log.Println("Error inserting users into database:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		user.Password = "" // Ensure password is not returned in the response
		return c.Status(fiber.StatusCreated).JSON(user)
	}
}
