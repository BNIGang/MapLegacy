package web

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func GenerateJWT(username string, secret []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTMiddleware(secret []byte, engine *html.Engine) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookie := c.Cookies("token")

		if cookie == "" {
			return c.Render("login", fiber.Map{"Error": "Incorrect username or password"})
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing token cookie"})
		}

		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
		if err != nil {
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
			return c.Render("login", fiber.Map{"Error": "Incorrect username or password"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("userID", claims["userID"])
			return c.Next()
		}

		// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
		return c.Render("login", fiber.Map{"Error": "Incorrect username or password"})
	}
}
