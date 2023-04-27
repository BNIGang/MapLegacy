package web

import (
	"github.com/gofiber/fiber/v2"
)

func GetCabang(c *fiber.Ctx) error {
	// Connect to your database
	db, err := Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	// Query the database to get bidang usaha data
	rows, err := db.Query("SELECT cabang_id, cabang_name FROM cabang ORDER BY cabang_name ASC")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer rows.Close()

	// Create a slice to store the bidang usaha data
	cabang := make([]struct {
		Cabang_ID   string `json:"cabang_id"`
		Cabang_Name string `json:"cabang_name"`
	}, 0)

	// Loop through the rows and add the data to the slice
	for rows.Next() {
		var cabang_id string
		var cabang_name string
		err = rows.Scan(&cabang_id, &cabang_name)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		cabang = append(cabang, struct {
			Cabang_ID   string `json:"cabang_id"`
			Cabang_Name string `json:"cabang_name"`
		}{Cabang_ID: cabang_id, Cabang_Name: cabang_name})
	}

	// Check for errors in the loop
	if err = rows.Err(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the bidang usaha data as a JSON response
	return c.JSON(cabang)
}

func GetKotaKabupatenHandler(c *fiber.Ctx) error {
	// Extract the cabang_id parameter from the request URL
	cabangID := c.Params("cabang_id")

	db, err := Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		SELECT kota_kabupaten_id, kota_kabupaten_name FROM kota_kabupaten WHERE cabang_id = ? ORDER BY kota_kabupaten_name ASC
	`)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer stmt.Close()

	rows, err := stmt.Query(cabangID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Create a slice to store the bidang usaha data
	kotaKabupaten := make([]struct {
		Kota_Kabupaten_ID   string `json:"kota_kabupaten_id"`
		Kota_Kabupaten_Name string `json:"kota_kabupaten_name"`
	}, 0)

	// Loop through the rows and add the data to the slice
	for rows.Next() {
		var kota_kabupaten_id string
		var kota_kabupaten_name string
		err = rows.Scan(&kota_kabupaten_id, &kota_kabupaten_name)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		kotaKabupaten = append(kotaKabupaten, struct {
			Kota_Kabupaten_ID   string `json:"kota_kabupaten_id"`
			Kota_Kabupaten_Name string `json:"kota_kabupaten_name"`
		}{Kota_Kabupaten_ID: kota_kabupaten_id, Kota_Kabupaten_Name: kota_kabupaten_name})
	}

	// Check for errors in the loop
	if err = rows.Err(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the bidang usaha data as a JSON response
	return c.JSON(kotaKabupaten)
}

func GetKCPKCUKK(c *fiber.Ctx) error {
	// Extract the cabang_id parameter from the request URL
	cabangID := c.Params("cabang_id")

	db, err := Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		SELECT kantor_id, kantor FROM KCU_KCP_KK WHERE cabang_id = ? ORDER BY kantor ASC
	`)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer stmt.Close()

	rows, err := stmt.Query(cabangID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Create a slice to store the bidang usaha data
	kcukcpkk := make([]struct {
		Kantor_ID string `json:"kantor_id"`
		Kantor    string `json:"kantor"`
	}, 0)

	// Loop through the rows and add the data to the slice
	for rows.Next() {
		var kantor_id string
		var kantor string
		err = rows.Scan(&kantor_id, &kantor)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		kcukcpkk = append(kcukcpkk, struct {
			Kantor_ID string `json:"kantor_id"`
			Kantor    string `json:"kantor"`
		}{Kantor_ID: kantor_id, Kantor: kantor})
	}

	// Check for errors in the loop
	if err = rows.Err(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the bidang usaha data as a JSON response
	return c.JSON(kcukcpkk)
}

func GetBidangUsahaHandler(c *fiber.Ctx) error {
	// Connect to your database
	db, err := Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	// Query the database to get bidang usaha data
	rows, err := db.Query("SELECT bidang_id, bidang FROM bidang_usaha ORDER BY bidang ASC")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer rows.Close()

	// Create a slice to store the bidang usaha data
	bidangUsaha := make([]struct {
		Bidang_ID string `json:"bidang_id"`
		Bidang    string `json:"bidang"`
	}, 0)

	// Loop through the rows and add the data to the slice
	for rows.Next() {
		var bidang_id string
		var bidang string
		err = rows.Scan(&bidang_id, &bidang)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		bidangUsaha = append(bidangUsaha, struct {
			Bidang_ID string `json:"bidang_id"`
			Bidang    string `json:"bidang"`
		}{Bidang_ID: bidang_id, Bidang: bidang})
	}

	// Check for errors in the loop
	if err = rows.Err(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the bidang usaha data as a JSON response
	return c.JSON(bidangUsaha)
}

func GetProdukUsahaHandler(c *fiber.Ctx) error {
	// Extract the bidang_id parameter from the request URL
	bidangID := c.Params("bidang_id")

	db, err := Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		SELECT produk_id, usaha FROM produk_usaha WHERE bidang_usaha_id = ? ORDER BY usaha ASC
	`)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer stmt.Close()

	rows, err := stmt.Query(bidangID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Create a slice to store the bidang usaha data
	produkUsaha := make([]struct {
		Produk_ID string `json:"produk_id"`
		Usaha     string `json:"usaha"`
	}, 0)

	// Loop through the rows and add the data to the slice
	for rows.Next() {
		var produk_id string
		var usaha string
		err = rows.Scan(&produk_id, &usaha)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		produkUsaha = append(produkUsaha, struct {
			Produk_ID string `json:"produk_id"`
			Usaha     string `json:"usaha"`
		}{Produk_ID: produk_id, Usaha: usaha})
	}

	// Check for errors in the loop
	if err = rows.Err(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the bidang usaha data as a JSON response
	return c.JSON(produkUsaha)
}
