package v1

import "fmt"

func MapLegacyHandler(nasabah_id string) {

	data_nasabah, err := GetNasabahByID(nasabah_id)
	if err != nil {
		return
	}

	for _, afiliasi := range data_nasabah.AfiliasiList {
		namaAfiliasi := afiliasi.NamaAfiliasi
		hubunganAfiliasi := afiliasi.HubunganAfiliasi

		// do something i dont know Lmao!
		fmt.Println(namaAfiliasi)
		fmt.Println(hubunganAfiliasi)
	}
}

// {
// 	nama: Lmao,
// 	afiliasi_list: {
// 		{
// 			nama: Lmao Jr.,
// 			afiliasi: anak,
// 			afiliasi_list: {
// 				{}
// 			}
// 		},
// 		{
// 			nama: Lmao Sr.,
// 			afiliasi: ayah,
// 			afiliasi_list: {
// 				nama: Mrs. Lmao,
// 				afiliasi: istri,
// 				afiliasi_list:{
// 					{}
// 				}
// 			}
// 		},
// 	}
// }

// {
// 	nama: Mrs.Lmao,
// 	afiliasi_list: {
// 		nama: Lmao Sr.,
// 		afiliasi: suami,
// 		afiliasi_list: {
// 			nama: Lmao,
// 			afiliasi: anak,
// 			afiliasi_list: {
// 				nama: Lmao Jr.,
// 				afiliasi: cucu,
// 				afiliasi_list: {
// 					{}
// 				}
// 			}

// 		}
// 	}
// }

// Lmao Sr., Mrs.Lmao, Lmao, Lmao Jr.,
