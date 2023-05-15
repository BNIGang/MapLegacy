package v1

type Node struct {
	Name         string  `json:"nama"`
	Afiliasi     string  `json:"afiliasi"`
	AfiliasiList []*Node `json:"afiliasi_list,omitempty"`
}

func GenerateHierarchy(data *Nasabah) *Node {
	root := &Node{
		Name:     data.Nama_pengusaha,
		Afiliasi: "", // Assuming the root node has no afiliasi
	}

	// Recursive function to build the hierarchy
	var buildHierarchy func(parentNode *Node, afiliasiList []Afiliasi)
	buildHierarchy = func(parentNode *Node, afiliasiList []Afiliasi) {
		for _, afiliasi := range afiliasiList {
			node := &Node{
				Name:     afiliasi.NamaAfiliasi,
				Afiliasi: afiliasi.HubunganAfiliasi,
			}
			parentNode.AfiliasiList = append(parentNode.AfiliasiList, node)
			buildHierarchy(node, afiliasi.AfiliasiList)
		}
	}

	buildHierarchy(root, data.AfiliasiList)
	return root
}

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
