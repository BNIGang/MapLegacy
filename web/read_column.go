package web

type KotaData struct {
	KotaKabupatenID   string
	KotaKabupatenName string
}

func GetKotaKabupaten(cabang_id string) ([]KotaData, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		SELECT kota_kabupaten_id, kota_kabupaten_name FROM kota_kabupaten WHERE cabang_id = ?
	`)
	if err != nil {
		return nil, err // database error
	}
	defer stmt.Close()

	rows, err := stmt.Query(cabang_id)
	if err != nil {
		return nil, err // database error
	}

	var kota []KotaData
	for rows.Next() {
		var kotaData KotaData
		err = rows.Scan(
			&kotaData.KotaKabupatenID,
			&kotaData.KotaKabupatenName,
		)
		if err != nil {
			return nil, err // database error
		}
		kota = append(kota, kotaData)
	}

	if err = rows.Err(); err != nil {
		return nil, err // database error
	}

	return kota, nil
}
