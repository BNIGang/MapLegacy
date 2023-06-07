package v1

import (
	"github.com/BNIGang/MapLegacy/login"
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

		// Prepare the statement to update the data
		stmt, err := db.Prepare(`
			UPDATE 
				users
			SET 
				name = ?,
				username = ?
			WHERE
				user_id = ?
		`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Check if password is not empty and matches confirm password
		if password != "" {
			if password != confirmpassword {
				alert := "Konfirmasi password tidak sama"
				return c.Redirect("/edit_user/" + user_id + "?alert=" + alert)
			}

			// Generate the hashed password
			hashedPassword, err := login.GenerateHash(password)
			if err != nil {
				return err
			}

			// Update the password field in the users table
			stmt.Exec(name, username, hashedPassword, user_id)
		} else {
			// Update the name and username fields in the users table without changing the password
			stmt.Exec(name, username, user_id)
		}

		// cabang_id, err := GetCabangID(cabang)
		// if err != nil {
		// 	return err
		// }

		// Prepare the statement to update the data
		stmt2, err := db.Prepare(`
			UPDATE 
				user_privileges
			SET 
				cabang_id = ?,
				user_privilege = ?
			WHERE 
				user_id = ?
		`)
		if err != nil {
			return err
		}
		defer stmt2.Close()

		// Execute the statement to update the data in the user_privileges table
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
