package login

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// Check if the username and password are correct (use placeholder values for now)
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "user" && password == "pass" {
		// Set a cookie to indicate that the user is logged in
		cookie := &fiber.Cookie{
			Name:  "token",
			Value: "jwt-token",
			Path:  "/",
		}
		c.Cookie(cookie)

		// Redirect to the home page
		return c.Redirect("/home", fiber.StatusSeeOther)
	}

	// If the username and password are incorrect, render the login page again with an error message
	return c.Render("login", fiber.Map{"Error": "Incorrect username or password"})
}
