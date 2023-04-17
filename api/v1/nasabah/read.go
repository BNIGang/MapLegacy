package v1

import (
	"database/sql"

	"github.com/BNIGang/MapLegacy/web"
)

type Nasabah struct {
	id                       string
	nama_pengusaha           string
	nomor_kontak             string
	alamat_tempat_tinggal    string
	bidang_usaha             string
	produk_usaha             string
	detail_bidang_usaha      string
	kabupaten_kota           string
	cabang                   string
	KCU_KCP_KK               string
	nasabah                  string
	no_CIF                   string
	AUM_di_BNI               string
	debitur                  string
	kredit_di_bni            string
	produk_bni_yang_dimiliki string
	mitra_bank_dominan       string
	aum_di_bank_lain         string
	kredit_di_bank_lain      string
	afiliasi                 string
	hubungan_afiliasi        string
	added_by                 string
}

func GetNasabahData(user_id string, wilayah_id string, cabang_id string, privilege string) (*Nasabah, error) {
	db, err := web.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// row := db.QueryRow(`
	// 	SELECT * FROM data_nasabah WHERE ???????`,  // Handle this
	// )

	row := db.QueryRow(`
		SELECT * FROM data_nasabah`, // Placeholder to test
	)

	var nasabah Nasabah
	err = row.Scan(
		&nasabah.id,
		&nasabah.nama_pengusaha,
		&nasabah.nomor_kontak,
		&nasabah.alamat_tempat_tinggal,
		&nasabah.bidang_usaha,
		&nasabah.produk_usaha,
		&nasabah.detail_bidang_usaha,
		&nasabah.kabupaten_kota,
		&nasabah.cabang,
		&nasabah.KCU_KCP_KK,
		&nasabah.nasabah,
		&nasabah.no_CIF,
		&nasabah.AUM_di_BNI,
		&nasabah.debitur,
		&nasabah.kredit_di_bni,
		&nasabah.produk_bni_yang_dimiliki,
		&nasabah.mitra_bank_dominan,
		&nasabah.aum_di_bank_lain,
		&nasabah.kredit_di_bank_lain,
		&nasabah.afiliasi,
		&nasabah.hubungan_afiliasi,
		&nasabah.added_by,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // nasabah not found
		}
		return nil, err // database error
	}

	return &nasabah, nil
}

func SearchNasabah(query string) {

}
