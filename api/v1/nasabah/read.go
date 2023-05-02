package v1

import (
	"fmt"

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
	Afiliasi                 string
	Hubungan_afiliasi        string
	Added_by                 string
	Username                 string
}

var nasabahMap map[string]Nasabah

func GetNasabahDataByUser(user_id string, wilayah_id string, cabang_id string, privilege string) ([]Nasabah, error) {
	nasabahMap = make(map[string]Nasabah)

	db, err := web.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// row := db.QueryRow(`
	// 	SELECT * FROM data_nasabah WHERE ???????`,  // Handle this to show correct data_nasabah for each privilege
	// )

	// rows, err := db.Query(`
	// 	SELECT * FROM data_nasabah
	// `)

	rows, err := db.Query(`
		SELECT dn.*, u.username
		FROM data_nasabah dn
		INNER JOIN users u ON dn.added_by = u.user_id
	`)

	var nasabahs []Nasabah
	for rows.Next() {
		var nasabah Nasabah
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
			&nasabah.Afiliasi,
			&nasabah.Hubungan_afiliasi,
			&nasabah.Added_by,
			&nasabah.Username,
		)
		if err != nil {
			return nil, err // database error
		}
		nasabahMap[nasabah.Id] = nasabah
		nasabahs = append(nasabahs, nasabah)
	}

	if err = rows.Err(); err != nil {
		return nil, err // database error
	}

	return nasabahs, nil
}

func GetNasabahByID(nasabah_id string) (Nasabah, error) {
	nasabah, ok := nasabahMap[nasabah_id]
	if !ok {
		return Nasabah{}, fmt.Errorf("nasabah with id %s not found", nasabah_id)
	}
	return nasabah, nil
}

func SearchNasabah(query string) {
	//TODO
}
