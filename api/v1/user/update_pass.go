package v1

import (
	"github.com/BNIGang/MapLegacy/login"
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func EditPassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the form data
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		// Retrieve the text fields
		userid := form.Value["userid"][0]
		oldpass := form.Value["oldpass"][0]
		newpass := form.Value["newpass"][0]
		confirmpass := form.Value["confirmpass"][0]

		//See if old pass match with one on database
		check := CheckOldPass(oldpass, userid)

		if !check {
			alert := "Password lama tidak sama!"
			return c.Redirect("/edit_password?alert=" + alert)
		}

		if newpass != confirmpass {
			alert := "Konfirmasi password tidak sama!"
			return c.Redirect("/edit_password?alert=" + alert)
		}

		if newpass == oldpass {
			alert := "Password baru tidak boleh sama dengan passwod lama!"
			return c.Redirect("/edit_password?alert=" + alert)
		}

		//If reached here, then password passes
		hashed_pass, err := login.GenerateHash(newpass)
		if err != nil {
			return err
		}

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
					password = ?
				WHERE 
					user_id = ?
				`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Execute the statement with the provided data and nasabah_id
		_, err = stmt.Exec(
			hashed_pass,
			userid,
		)
		if err != nil {
			return err
		}

		// Give success prompt here
		alert := "Password berhasil dirubah!"
		return c.Redirect("/edit_password?alert=" + alert)
	}
}

func CheckOldPass(password string, userid string) bool {
	db, err := web.Connect()
	if err != nil {
		return false
	}
	defer db.Close()

	var oldpass string
	err = db.QueryRow("SELECT password FROM users WHERE user_id = ?", userid).Scan(&oldpass)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(oldpass), []byte(password))
	if err != nil {
		return false
	}

	return true
}
