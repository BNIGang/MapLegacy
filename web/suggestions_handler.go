package web

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Suggestion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func AutoFillHandler(c *fiber.Ctx) error {
	nama_pengusaha := c.Params("nama_pengusaha")

	db, err := Connect()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	// query := `
	// 	SELECT
	// 		a.id_parent, dn.nama_pengusaha
	// 	FROM
	// 		afiliasi a
	// 	LEFT JOIN
	// 		data_nasabah dn
	// 	ON
	// 		a.id_parent = dn.id
	// 	WHERE
	// 		dn.nama_pengusaha LIKE ?
	// `

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
		combined_column LIKE ?;
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
