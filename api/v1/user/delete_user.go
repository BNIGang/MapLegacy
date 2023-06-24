package v1

import (
	"log"

	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
)

func DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the cabang_id parameter from the request URL
		user_id := c.Params("user_id")

		db, err := web.Connect()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer db.Close()

		stmt, err := db.Prepare("DELETE FROM users WHERE user_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(user_id)
		if err != nil {
			log.Fatal(err)
		}

		stmt, err = db.Prepare("DELETE FROM user_privileges WHERE user_id=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(user_id)
		if err != nil {
			log.Fatal(err)
		}

		// Redirect to home page
		return c.Redirect("/user_page")
	}
}
