package main

import (
	"fmt"

	v1 "github.com/BNIGang/MapLegacy/api/v1/nasabah"
	"github.com/BNIGang/MapLegacy/login"
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/golang-jwt/jwt"
)

// Change this later
var secret []byte = []byte("super-secret-key")

func main() {

	engine := html.New("./web/template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Login page
	app.Get("/", func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")
		if cookie != "" {
			return c.Redirect("/home")
		}
		return c.Render("login", fiber.Map{})
	})

	app.Post("/login", login.Handler(engine))

	app.Get("/home", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		// get Username from cookie
		// cookie contains JWT token, decrypt the token to get username
		cookie := c.Cookies("token")
		var username string

		if cookie == "" {
			return c.Render("login", fiber.Map{"Error": "Missing token cookie"})
		}

		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, c.Render("login", fiber.Map{"Error": "Unexpected Signing Method"})
			}
			return secret, nil
		})

		if err != nil {
			return c.Render("login", fiber.Map{"Error": fmt.Sprintf("Invalid or expired token: %v", err)})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if val, ok := claims["username"].(string); ok {
				username = val
			}
		}

		user, err := web.GetUserByUsername(username)
		if err != nil {
			return c.Render("login", fiber.Map{"Error": err})
		}
		if user == nil {
			return c.Render("login", fiber.Map{"Error": err})
		}

		data_nasabah, err := v1.GetNasabahData()

		return c.Render("home", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
		})
	})

	app.Get("/home/", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return c.Redirect("/home")
	})

	app.Get("/logout", login.LogoutHandler)

	port := ":8000"
	app.Listen(port)
}