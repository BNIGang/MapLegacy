package web

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/golang-jwt/jwt"
)

// Still use cookie for now, make proper auth later for command line users
func GenerateJWT(username string, secret []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix() // expires in 1 hour
	claims["refreshed"] = false

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
			return c.Render("login", fiber.Map{"Error": "Mising token"})
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing token cookie"})
		}

		tokenString, err := RefreshJWT(c, secret, engine)
		if err != nil {
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
			return c.Render("login", fiber.Map{"Error": "Invalid or expired token"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
		if err != nil {
			return c.Render("login", fiber.Map{"Error": "Invalid or expired token"})
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("username", claims["username"])
			return c.Next()
		}

		return c.Render("login", fiber.Map{"Error": "Invalid or expired token"})
	}
}

func RefreshJWT(c *fiber.Ctx, secret []byte, engine *html.Engine) (string, error) {
	cookie := c.Cookies("token")
	if cookie == "" {
		return "", c.Render("login", fiber.Map{"Error": "Missing token cookie"})
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return "", c.Render("login", fiber.Map{"Error": fmt.Sprintf("Invalid or expired token: %v", err)})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Until(expirationTime) > 5*time.Minute {
			return cookie, nil
		}

		newToken := jwt.New(jwt.SigningMethodHS256)
		newClaims := newToken.Claims.(jwt.MapClaims)

		for key, value := range claims {
			newClaims[key] = value
		}

		newClaims["exp"] = time.Now().Add(time.Hour).Unix()
		newClaims["refreshed"] = true

		newTokenString, err := newToken.SignedString(secret)
		if err != nil {
			return "", c.Render("login", fiber.Map{"Error": fmt.Sprintf("Error generating new token: %v", err)})
		}

		newCookie := &fiber.Cookie{
			Name:  "token",
			Value: newTokenString,
			Path:  "/",
		}
		c.Cookie(newCookie)

		return newTokenString, nil
	}

	return "", c.Render("login", fiber.Map{"Error": "Invalid or expired token"})
}
