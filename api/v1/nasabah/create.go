package v1

import (
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
)

func AddNasabahHandler(user_id string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the form data
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		// Retrieve the text fields
		pengusaha := form.Value["nama_pengusaha"][0]
		nomor_kontak := form.Value["nomor_kontak"][0]
		alamat_tempat_tinggal := form.Value["alamat_tempat_tinggal"][0]
		bidang_usaha := form.Value["bidang_usaha"][0]
		produk_usaha := form.Value["produk_usaha"][0]
		detail_bidang_usaha := form.Value["detail_bidang_usaha"][0]
		kabupaten_kota := form.Value["kabupaten_kota"][0]
		cabang := form.Value["cabang"][0]
		kcu_kcp_kk := form.Value["kcu_kcp_kk"][0]
		nasabah := form.Value["nasabah"][0]
		no_CIF := form.Value["no_CIF"][0]
		aum_di_bni := form.Value["aum_di_bni"][0]
		debitur := form.Value["debitur"][0]
		kredit_di_bni := form.Value["kredit_di_bni"][0]
		produk_bni_yang_dimiliki := form.Value["produk_bni_yang_dimiliki"][0]
		mitra_bank_dominan := form.Value["mitra_bank_dominan"][0]
		aum_di_bank_lain := form.Value["aum_di_bank_lain"][0]
		kredit_di_bank_lain := form.Value["kredit_di_bank_lain"][0]
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
			user_id,
		)
		if err != nil {
			return err
		}

		// This part to add data afiliasi to afiliasi table
		stmt2, err2 := db.Prepare(`
		INSERT INTO afiliasi 
		(
			id_child,
			id_parent,
			nama_child,
			hubungan
		) VALUES 
		(
			UUID(),
			(SELECT id FROM data_nasabah WHERE nama_pengusaha = ?),
			?,
			?
		)
		`)
		if err2 != nil {
			return err
		}
		defer stmt2.Close()

		// Retrieve the array values
		afiliasiValues := form.Value["afiliasi[]"]
		hubunganAfiliasiValues := form.Value["hubungan_afiliasi[]"]

		// Iterate over the array values and process them accordingly
		for i := 0; i < len(afiliasiValues); i++ {
			afiliasi := afiliasiValues[i]
			hubunganAfiliasi := hubunganAfiliasiValues[i]

			// Execute the SQL statement with the current values
			_, err := stmt2.Exec(pengusaha, afiliasi, hubunganAfiliasi)
			if err != nil {
				return err
			}
		}

		// Redirect to home page
		return c.Redirect("/home")
	}
}

func AddAfiliasi(user_id string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		db, err := web.Connect()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer db.Close()

		// This part to add data afiliasi to afiliasi table
		stmt, err := db.Prepare(`
		INSERT INTO afiliasi 
		(
			id_child,
			id_parent,
			nama_child,
			hubungan,
			added_by
		) VALUES 
		(
			UUID(),
			?,
			?,
			?,
			?
		)
		`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Retrieve the text fields
		// pengusaha := form.Value["nama_pengusaha"][0]
		id_pengusaha := form.Value["id_pengusaha"][0]

		// Retrieve the array values
		afiliasiValues := form.Value["afiliasi[]"]
		hubunganAfiliasiValues := form.Value["hubungan_afiliasi[]"]

		// Iterate over the array values and process them accordingly
		for i := 0; i < len(afiliasiValues); i++ {
			afiliasi := afiliasiValues[i]
			hubunganAfiliasi := hubunganAfiliasiValues[i]

			// Execute the SQL statement with the current values
			_, err := stmt.Exec(id_pengusaha, afiliasi, hubunganAfiliasi, user_id)
			if err != nil {
				return err
			}
		}

		return c.Redirect("/afiliasi")
	}
}
