package web

import (
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Suggestion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func AutoFillHandler(c *fiber.Ctx) error {
	nama_pengusaha := c.Params("nama_pengusaha")

	// Decode the URL-encoded query parameter
	nama_pengusaha, err := url.QueryUnescape(nama_pengusaha)
	if err != nil {
		// Handle the error, e.g., return an error response
		return err
	}

	db, err := Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	query := `
	SELECT * FROM (
		SELECT 
			id_child, nama_child 
		AS 
			combined_column 
		FROM 
			afiliasi
		UNION
		SELECT 
			id, nama_pengusaha 
		FROM 
			data_nasabah
	) AS subquery
	WHERE 
		combined_column LIKE ?
	LIMIT
		5;
	`

	rows, err := db.Query(query, "%"+nama_pengusaha+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	suggestions := []Suggestion{}
	for rows.Next() {
		var idParent, namaPengusaha string
		if err := rows.Scan(&idParent, &namaPengusaha); err != nil {
			log.Fatal(err)
		}

		suggestion := Suggestion{
			ID:   idParent,
			Name: namaPengusaha,
		}
		suggestions = append(suggestions, suggestion)
	}

	response := map[string]interface{}{
		"suggestions": suggestions,
	}

	return c.JSON(response)
}
