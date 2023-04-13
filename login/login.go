package login

import (
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Handler(engine *html.Engine) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the username and password are correct (use placeholder values for now)
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username == "user" && password == "pass" {
			token, err := web.GenerateJWT(username, []byte("super-secret-key"))
			if err != nil {
				return err
			}

			cookie := &fiber.Cookie{
				Name:  "token",
				Value: token,
				Path:  "/",
			}
			c.Cookie(cookie)

			return c.Render("home", fiber.Map{"Name": token})
		}

		// If the username and password are incorrect, render the login page again with an error message
		return c.Render("login", fiber.Map{"Error": "Incorrect username or password"})
	}
}
