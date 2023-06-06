package v1

import (
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(user_id string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the form data
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		// Retrieve the text fields
		name := form.Value["name"][0]
		username := form.Value["username"][0]
		password := form.Value["password"][0]
		confirmpassword := form.Value["confirmpassword"][0]
		cabang := form.Value["cabang"][0]
		user_privilege := form.Value["user_privilege"][0]

		db, err := web.Connect()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer db.Close()

		//TODO: do tihs properly
		if password != "" {
			if password != confirmpassword {
				return c.Redirect("/home")
			}
		}

		// Prepare the statement to update the data
		stmt, err := db.Prepare(`
				UPDATE 
					users
				SET 
					name = ?,
					username = ?,
					password = ?,
				WHERE 
					id = ?
				`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Execute the statement with the provided data and nasabah_id
		_, err = stmt.Exec(
			name,
			username,
			password,
			user_id,
		)
		if err != nil {
			return err
		}

		// Prepare the statement to update the data
		stmt2, err := db.Prepare(`
				UPDATE 
					user_privileges
				SET 
					cabang_id = ?,
					user_privilege = ?,
				WHERE 
					id = ?
				`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Execute the statement with the provided data and nasabah_id
		_, err = stmt2.Exec(
			cabang,
			user_privilege,
			user_id,
		)
		if err != nil {
			return err
		}

		return c.Redirect("/user_page")
	}
}
