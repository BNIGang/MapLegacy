package login

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(c *fiber.Ctx) error {
	// Create a new cookie with an empty value and an expiration time in the past
	cookie := &fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	}

	c.Cookie(cookie)

	return c.Redirect("/")
}
