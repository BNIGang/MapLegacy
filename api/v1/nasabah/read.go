package v1

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/BNIGang/MapLegacy/web"
)

type Nasabah struct {
	Id                       string
	Nama_pengusaha           string
	Nomor_kontak             string
	Alamat_tempat_tinggal    string
	Bidang_usaha             string
	Produk_usaha             string
	Detail_bidang_usaha      string
	Kabupaten_kota           string
	Cabang                   string
	KCU_KCP_KK               string
	Nasabah                  string
	No_CIF                   string
	AUM_di_BNI               string
	Debitur                  string
	Kredit_di_bni            string
	Produk_bni_yang_dimiliki string
	Mitra_bank_dominan       string
	Aum_di_bank_lain         string
	Kredit_di_bank_lain      string
	Added_by                 string
	Username                 string
	AfiliasiList             []Afiliasi
}

type Afiliasi struct {
	IdChild          string
	IdParent         string
	NamaAfiliasi     string
	HubunganAfiliasi string
	AddedBy          string
	NamaPengusaha    string
	Username         string
}

type MergedRow struct {
	NamaPengusaha  string
	MergedAfiliasi []Afiliasi
	RowCount       int
}

var nasabahMap = make(map[string]*Nasabah)
var mergedMap = make(map[string]MergedRow)
var idChildMap = make(map[string]Afiliasi)
var mergedRows = make([]MergedRow, 0, len(mergedMap))

func GetNasabahDataByUser(user_id string, wilayah_id string, cabang_id string, privilege string) ([]Nasabah, error) {
	nasabahMap = make(map[string]*Nasabah)
	idChildMap = make(map[string]Afiliasi)

	db, err := web.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// TODO, query only based on wilayah/cabang/privilege/etc
	var query string
	var args []interface{}

	if privilege == "pemimpin_cabang" || privilege == "pemimpin_cabang_pembantu" {
		// Retrieve cabang_name based on cabang_id
		var cabangName string
		err := db.QueryRow("SELECT cabang_name FROM cabang WHERE cabang_id=?", cabang_id).Scan(&cabangName)
		if err != nil {
			return nil, err
		}

		// Set the WHERE clause in the query
		query = `
		SELECT 
			dn.*, 
			GROUP_CONCAT(a.nama_child) AS nama_child, 
			GROUP_CONCAT(a.hubungan) AS hubungan, 
			u.name 
		FROM 
			data_nasabah dn 
		LEFT JOIN 
			afiliasi a 
		ON  
			dn.id = a.id_parent 
		LEFT JOIN 
			users u 
		ON 
			dn.added_by = u.user_id 
		WHERE 
			dn.cabang = ? 
		GROUP BY 
			dn.id 
		ORDER BY 
			dn.nama_pengusaha ASC
	`
		args = append(args, cabangName)
	} else if privilege == "individu" {
		// Retrieve username based on user_id
		var name string
		err := db.QueryRow("SELECT name FROM users WHERE user_id=?", user_id).Scan(&name)
		if err != nil {
			return nil, err
		}

		// Set the WHERE clause in the query
		query = `
		SELECT 
			dn.*, 
			GROUP_CONCAT(a.nama_child) AS nama_child, 
			GROUP_CONCAT(a.hubungan) AS hubungan, 
			u.name 
		FROM 
			data_nasabah dn 
		LEFT JOIN 
			afiliasi a 
		ON  
			dn.id = a.id_parent 
		LEFT JOIN 
			users u 
		ON 
			dn.added_by = u.user_id 
		WHERE 
			u.name = ? 
		GROUP BY 
			dn.id 
		ORDER BY 
			dn.nama_pengusaha ASC
	`
		args = append(args, name)
	} else {
		// No additional WHERE clause
		query = `
		SELECT 
			dn.*, 
			GROUP_CONCAT(a.nama_child) AS nama_child, 
			GROUP_CONCAT(a.hubungan) AS hubungan, 
			u.name 
		FROM 
			data_nasabah dn 
		LEFT JOIN 
			afiliasi a 
		ON  
			dn.id = a.id_parent 
		LEFT JOIN 
			users u 
		ON 
			dn.added_by = u.user_id 
		GROUP BY 
			dn.id 
		ORDER BY 
			dn.nama_pengusaha ASC
	`
	}

	// Execute the query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nasabah Nasabah
		var afiliasi, hubunganAfiliasi sql.NullString

		err = rows.Scan(
			&nasabah.Id,
			&nasabah.Nama_pengusaha,
			&nasabah.Nomor_kontak,
			&nasabah.Alamat_tempat_tinggal,
			&nasabah.Bidang_usaha,
			&nasabah.Produk_usaha,
			&nasabah.Detail_bidang_usaha,
			&nasabah.Kabupaten_kota,
			&nasabah.Cabang,
			&nasabah.KCU_KCP_KK,
			&nasabah.Nasabah,
			&nasabah.No_CIF,
			&nasabah.AUM_di_BNI,
			&nasabah.Debitur,
			&nasabah.Kredit_di_bni,
			&nasabah.Produk_bni_yang_dimiliki,
			&nasabah.Mitra_bank_dominan,
			&nasabah.Aum_di_bank_lain,
			&nasabah.Kredit_di_bank_lain,
			&nasabah.Added_by,
			&afiliasi,
			&hubunganAfiliasi,
			&nasabah.Username,
		)
		if err != nil {
			return nil, err // database error
		}

		// Check if the nasabah is already in the map
		if _, ok := nasabahMap[nasabah.Id]; !ok {
			// If not, add it to the map with an empty list of afiliasi
			nasabah.AfiliasiList = make([]Afiliasi, 0)
			nasabahMap[nasabah.Id] = &nasabah
		}

		// If the afiliasi is not null, add it to the nasabah's list of afiliasi
		if afiliasi.Valid {
			afiliasiSlice := strings.Split(afiliasi.String, ",")
			hubunganAfiliasiSlice := strings.Split(hubunganAfiliasi.String, ",")
			for i := range afiliasiSlice {
				// Check if the afiliasi is already in the nasabah's list
				alreadyExists := false
				for j := range nasabahMap[nasabah.Id].AfiliasiList {
					if nasabahMap[nasabah.Id].AfiliasiList[j].NamaAfiliasi == afiliasiSlice[i] {
						alreadyExists = true
						break
					}
				}
				// If the afiliasi is not already in the nasabah's list, add it
				if !alreadyExists {
					nasabahMap[nasabah.Id].AfiliasiList = append(nasabahMap[nasabah.Id].AfiliasiList, Afiliasi{
						NamaAfiliasi:     afiliasiSlice[i],
						HubunganAfiliasi: hubunganAfiliasiSlice[i],
					})
				}
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Create a slice of Nasabah from the values in the map
	nasabahs := make([]Nasabah, 0, len(nasabahMap))
	for _, nasabah := range nasabahMap {
		nasabahs = append(nasabahs, *nasabah)
	}

	sort.Slice(nasabahs, func(i, j int) bool {
		return nasabahs[i].Nama_pengusaha < nasabahs[j].Nama_pengusaha
	})

	return nasabahs, nil
}

func GetNasabahByID(nasabah_id string) (*Nasabah, error) {
	nasabah, ok := nasabahMap[nasabah_id]
	if !ok {
		return nil, fmt.Errorf("nasabah with id %s not found", nasabah_id)
	}
	return nasabah, nil
}

func GetAfiliasiByUser(user_id string, wilayah_id string, cabang_id string, privilege string) (map[string]MergedRow, error) {
	mergedMap = make(map[string]MergedRow)

	//TODO , another one of those wilayah only thing

	db, err := web.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT 
            a.*, dn.nama_pengusaha, u.name
        FROM
            afiliasi a
        LEFT JOIN
            data_nasabah dn
        ON
            a.id_parent = dn.id
        LEFT JOIN
            users u
        ON
            a.added_by = u.user_id
    `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var afiliasi Afiliasi

		err = rows.Scan(
			&afiliasi.IdChild,
			&afiliasi.IdParent,
			&afiliasi.NamaAfiliasi,
			&afiliasi.HubunganAfiliasi,
			&afiliasi.AddedBy,
			&afiliasi.NamaPengusaha,
			&afiliasi.Username,
		)
		if err != nil {
			return nil, err // database error
		}

		// Check if the NamaPengusaha is already in the mergedMap
		if _, ok := mergedMap[afiliasi.IdParent]; !ok {
			mergedMap[afiliasi.IdParent] = MergedRow{
				MergedAfiliasi: []Afiliasi{afiliasi},
				RowCount:       1,
			}
		} else {
			mergedRow := mergedMap[afiliasi.IdParent]
			mergedRow.MergedAfiliasi = append(mergedRow.MergedAfiliasi, afiliasi)
			mergedRow.RowCount++
			mergedMap[afiliasi.IdParent] = mergedRow
		}

		idChildMap[afiliasi.IdChild] = afiliasi

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return mergedMap, nil
}

func GetAfiliasiById(id_child string) (*Afiliasi, error) {
	afiliasi, ok := idChildMap[id_child]
	if !ok {
		return nil, fmt.Errorf("Afiliasi with id %s not found", id_child)
	}
	return &afiliasi, nil
}

func GetAfiliasiListById(id_parent string) (*MergedRow, error) {
	afiliasi, ok := mergedMap[id_parent]
	if !ok {
		return nil, fmt.Errorf("Afiliasi list with id %s not found", id_parent)
	}
	return &afiliasi, nil
}

// func SearchNasabah(query string) {
// 	//TODO
// }
