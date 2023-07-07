package v1

import (
	"log"

	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
)

func DeleteNasabahData(c *fiber.Ctx) error {
	// Extract the cabang_id parameter from the request URL
	nasabah_id := c.Params("nasabah_id")

	db, err := web.Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM data_nasabah WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(nasabah_id)
	if err2 != nil {
		log.Fatal(err)
	}

	stmt2, err3 := db.Prepare("DELETE FROM afiliasi WHERE id_parent=?")
	if err3 != nil {
		log.Fatal(err3)
	}
	defer stmt2.Close()

	_, err4 := stmt2.Exec(nasabah_id)
	if err4 != nil {
		log.Fatal(err)
	}

	// Redirect to home page
	return c.Redirect("/home")
}

func DeleteAfiliasiData(c *fiber.Ctx) error {
	// Extract the cabang_id parameter from the request URL
	afiliasi_id := c.Params("afiliasi_id")

	db, err := web.Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM afiliasi WHERE (id_child=? OR id_parent=?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(afiliasi_id, afiliasi_id)
	if err2 != nil {
		log.Fatal(err)
	}

	return c.Redirect("/afiliasi")
}
