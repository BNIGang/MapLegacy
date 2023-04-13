package main

import (
	"github.com/BNIGang/MapLegacy/login"
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var secret []byte = []byte("super-secret-key")

func main() {

	engine := html.New("./web/template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Login page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	app.Post("/login", login.Handler(engine))

	app.Get("/home", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		// TODO: add name
		return c.Render("home", fiber.Map{})
	})

	app.Get("/logout", login.LogoutHandler)

	port := ":8000"
	app.Listen(port)
}
