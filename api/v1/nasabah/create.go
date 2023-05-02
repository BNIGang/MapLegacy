package v1

import (
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
)

func AddNasabahHandler(user_id string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pengusaha := c.FormValue("nama_pengusaha")
		nomor_kontak := c.FormValue("nomor_kontak")
		alamat_tempat_tinggal := c.FormValue("alamat_tempat_tinggal")
		bidang_usaha := c.FormValue("bidang_usaha")
		produk_usaha := c.FormValue("produk_usaha")
		detail_bidang_usaha := c.FormValue("detail_bidang_usaha")
		kabupaten_kota := c.FormValue("kabupaten_kota")
		cabang := c.FormValue("cabang")
		kcu_kcp_kk := c.FormValue("kcu_kcp_kk")
		nasabah := c.FormValue("nasabah")
		no_CIF := c.FormValue("no_CIF")
		aum_di_bni := c.FormValue("aum_di_bni")
		debitur := c.FormValue("debitur")
		kredit_di_bni := c.FormValue("kredit_di_bni")
		produk_bni_yang_dimiliki := c.FormValue("produk_bni_yang_dimiliki")
		mitra_bank_dominan := c.FormValue("mitra_bank_dominan")
		aum_di_bank_lain := c.FormValue("aum_di_bank_lain")
		kredit_di_bank_lain := c.FormValue("kredit_di_bank_lain")
		afiliasi := c.FormValue("afiliasi")
		hubungan_afiliasi := c.FormValue("hubungan_afiliasi")
		user_id := user_id

		db, err := web.Connect()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer db.Close()

		// Prepare the statement to insert the data
		stmt, err := db.Prepare(`INSERT INTO data_nasabah 
							(
								id,
								nama_pengusaha,
								nomor_kontak,
								alamat_tempat_tinggal,
								bidang_usaha,
								produk_usaha,
								detail_bidang_usaha,
								kabupaten_kota,
								cabang,
								kcu_kcp_kk,
								nasabah,
								no_CIF,
								aum_di_bni,
								debitur,
								kredit_di_bni,
								produk_bni_yang_dimiliki,
								mitra_bank_dominan,
								aum_di_bank_lain,
								kredit_di_bank_lain,
								afiliasi,
								hubungan_afiliasi,
								added_by
							) VALUES (
								UUID(),
								?,
								?,
								?,
								(SELECT bidang FROM bidang_usaha WHERE bidang_id = ?),
								(SELECT usaha FROM produk_usaha WHERE produk_id = ?),
								?,
								(SELECT kota_kabupaten_name FROM kota_kabupaten WHERE kota_kabupaten_id = ?),
								(SELECT cabang_name FROM cabang WHERE cabang_id = ?),
								(SELECT kantor FROM kantor WHERE kantor_id = ?),
								?,
								?,
								?,
								?,
								?,
								?,
								?,
								?,
								?,
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
			pengusaha,
			nomor_kontak,
			alamat_tempat_tinggal,
			bidang_usaha,
			produk_usaha,
			detail_bidang_usaha,
			kabupaten_kota,
			cabang,
			kcu_kcp_kk,
			nasabah,
			no_CIF,
			aum_di_bni,
			debitur,
			kredit_di_bni,
			produk_bni_yang_dimiliki,
			mitra_bank_dominan,
			aum_di_bank_lain,
			kredit_di_bank_lain,
			afiliasi,
			hubungan_afiliasi,
			user_id,
		)
		if err != nil {
			return err
		}

		// Redirect to home page
		return c.Redirect("/home")
	}
}
