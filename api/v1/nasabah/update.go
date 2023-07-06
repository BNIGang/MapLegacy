package v1

import (
	"fmt"

	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
)

func UpdateNasabahData(user_id string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		nasabah_id := c.Params("nasabah_id")

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
		latitude := form.Value["latitude"][0]
		longtitude := form.Value["longtitude"][0]
		user_id := user_id

		db, err := web.Connect()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer db.Close()

		// Prepare the statement to update the data
		stmt, err := db.Prepare(`
				UPDATE 
					data_nasabah 
				SET 
					nama_pengusaha = ?,
					nomor_kontak = ?,
					alamat_tempat_tinggal = ?,
					bidang_usaha = (SELECT bidang FROM bidang_usaha WHERE bidang_id = ?),
					produk_usaha = (SELECT usaha FROM produk_usaha WHERE produk_id = ?),
					detail_bidang_usaha = ?,
					kabupaten_kota = (SELECT kota_kabupaten_name FROM kota_kabupaten WHERE kota_kabupaten_id = ?),
					cabang = (SELECT cabang_name FROM cabang WHERE cabang_id = ?),
					kcu_kcp_kk = (SELECT kantor FROM kantor WHERE kantor_id = ?),
					nasabah = ?,
					no_CIF = ?,
					aum_di_bni = ?,
					debitur = ?,
					kredit_di_bni = ?,
					produk_bni_yang_dimiliki = ?,
					mitra_bank_dominan = ?,
					aum_di_bank_lain = ?,
					kredit_di_bank_lain = ?,
					latitude = ?,
					longtitude = ?
				WHERE 
					id = ?
				`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Execute the statement with the provided data and nasabah_id
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
			latitude,
			longtitude,
			nasabah_id,
		)
		if err != nil {
			return err
		}

		afiliasiValues := form.Value["afiliasi[]"]
		hubunganAfiliasiValues := form.Value["hubungan_afiliasi[]"]
		originalAfiliasiValues := form.Value["original_afiliasi[]"]
		idChildValues := form.Value["id_child[]"]
		deletedAfiliasiValues := form.Value["deleted_afiliasi[]"]

		// Update or insert afiliasi data
		for i := 0; i < len(afiliasiValues); i++ {
			afiliasi := afiliasiValues[i]
			hubunganAfiliasi := hubunganAfiliasiValues[i]
			originalAfiliasi := originalAfiliasiValues[i]
			idChild := idChildValues[i]

			if afiliasi != "" {
				if originalAfiliasi == "" {
					// Run the INSERT query
					fmt.Println("reached here")

					stmt2, err2 := db.Prepare(`
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
							(SELECT id FROM data_nasabah WHERE nama_pengusaha = ?),
							?,
							?,
							?
						)
					`)
					if err2 != nil {
						return err2
					}
					defer stmt2.Close()

					_, err := stmt2.Exec(pengusaha, afiliasi, hubunganAfiliasi, user_id)
					if err != nil {
						return err
					}
				} else if originalAfiliasi != afiliasi {
					// Run the UPDATE query
					fmt.Println("reached here instead")
					stmt2, err2 := db.Prepare(`
						UPDATE afiliasi
						SET nama_child = ?, hubungan = ?
						WHERE id_child = ?
					`)
					if err2 != nil {
						return err2
					}
					defer stmt2.Close()

					_, err := stmt2.Exec(afiliasi, hubunganAfiliasi, idChild)
					if err != nil {
						return err
					}
				}
			}
		}

		// Delete afiliasi data
		for _, deletedAfiliasi := range deletedAfiliasiValues {
			stmtDelete, errDelete := db.Prepare(`
				DELETE FROM afiliasi WHERE id_child = ?
			`)
			if errDelete != nil {
				return errDelete
			}
			defer stmtDelete.Close()

			_, err := stmtDelete.Exec(deletedAfiliasi)
			if err != nil {
				return err
			}
		}

		return c.Redirect("/home")
	}
}

func UpdateAfiliasi(user_id string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		afiliasi_id := c.Params("afiliasi_id")

		// Retrieve the form data
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		// Retrieve the text fields
		afiliasi := form.Value["afiliasi[]"][0]
		hubunganAfiliasi := form.Value["hubungan_afiliasi[]"][0]

		db, err := web.Connect()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer db.Close()

		stmt, err := db.Prepare(`
				UPDATE 
					afiliasi
				SET
					nama_child = ?,
					hubungan = ?
				WHERE
					id_child = ?
				`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err2 := stmt.Exec(afiliasi, hubunganAfiliasi, afiliasi_id)
		if err2 != nil {
			return err2
		}

		return c.Redirect("/afiliasi")
	}
}
