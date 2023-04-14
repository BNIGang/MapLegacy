package login

import (
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Handler(engine *html.Engine) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if web.AuthenticateUser(username, password) {
			token, err := web.GenerateJWT(username, []byte("super-secret-key"))
			if err != nil {
				return c.Render("login", fiber.Map{"Error": err})
			}

			cookie := &fiber.Cookie{
				Name:  "token",
				Value: token,
				Path:  "/",
			}

			c.Cookie(cookie)

			// return c.Render("home", fiber.Map{"Name": username})
			return c.Redirect("/home")
		}

		return c.Render("login", fiber.Map{"Error": "Incorrect username or password"})
	}
}
