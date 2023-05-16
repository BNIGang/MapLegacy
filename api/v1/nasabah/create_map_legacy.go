package v1

func MapLegacyHandler(nasabah_id string) ([]Afiliasi, error) {
	data_nasabah, err := GetNasabahByID(nasabah_id)
	if err != nil {
		return nil, err
	}

	return data_nasabah.AfiliasiList, nil
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
