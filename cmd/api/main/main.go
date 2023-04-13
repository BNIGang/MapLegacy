package main

import (
	"github.com/BNIGang/MapLegacy/login"
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

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

	app.Get("/home", web.JWTMiddleware([]byte("super-secret-key"), engine), func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})

	port := ":8000"
	app.Listen(port)
}
