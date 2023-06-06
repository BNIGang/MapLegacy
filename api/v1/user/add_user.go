package v1

import (
	"github.com/BNIGang/MapLegacy/login"
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
)

func AddUsersHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the form data
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		// Retrieve the text fields
		nama_lengkap := form.Value["name"][0]
		username := form.Value["username"][0]
		password := form.Value["password"][0]
		confirmpassword := form.Value["confirmpassword"][0]
		cabang := form.Value["cabang"][0]
		privilege := form.Value["user_privilege"][0]

		//Check if user exist
		exists, err := userExists(username)
		if err != nil {
			return err // Handle the database error
		}

		// If user exists, render a page with the error message
		if exists {
			alert := "Username sudah ada, tolong pilih Username lain."
			return c.Redirect("/add_users?alert=" + alert)
		}

		// Pass does not match
		if confirmpassword != password {
			alert := "Konfirmasi password tidak sama"
			return c.Redirect("/add_users?alert=" + alert)
		}

		//hash password
		hashed_pass, err := login.GenerateHash(password)
		if err != nil {
			return err
		}

		db, err := web.Connect()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer db.Close()

		// Prepare the statement to insert the data
		stmt, err := db.Prepare(`INSERT INTO users 
							(
								user_id,
								name,
								username,
								password
							) VALUES (
								UUID(),
								?,
								?,
								?
							)`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Execute the statement with the provided data
		_, err = stmt.Exec(
			nama_lengkap,
			username,
			hashed_pass,
		)
		if err != nil {
			return err
		}

		// Prepare the statement to insert the data
		stmt2, err := db.Prepare(`INSERT INTO user_privileges
							(
								user_id,
								wilayah_id,
								cabang_id,
								user_privilege
							) VALUES (
								(SELECT user_id FROM users WHERE username=?),
								"1",
								?,
								?
							)`)
		if err != nil {
			return err
		}
		defer stmt2.Close()

		// Execute the statement with the provided data
		_, err = stmt2.Exec(
			username,
			cabang,
			privilege,
		)
		if err != nil {
			return err
		}

		// Redirect to home page
		return c.Redirect("/user_page")
	}
}

func userExists(username string) (bool, error) {
	db, err := web.Connect()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
