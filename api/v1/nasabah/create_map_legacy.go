package v1

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BNIGang/MapLegacy/web"
)

type Node struct {
	Child    string `json:"child"`
	Parent   string `json:"parent"`
	Afiliasi string `json:"afiliasi"`
}

func MapLegacyHandler(afiliasi *MergedRow) ([]Node, error) {
	db, err := web.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if len(afiliasi.MergedAfiliasi) == 0 {
		firstNode := Node{
			Child:    afiliasi.NamaPengusaha,
			Parent:   "",
			Afiliasi: "",
		}
		return []Node{firstNode}, nil
	}

	var nodes []Node
	nodeMap := make(map[string]int) // Map to store the index of existing nodes by child ID

	// Add the first node with Child as a.IdParent and Parent as an empty string
	rootName, err := TranslateIdIntoName(db, afiliasi.MergedAfiliasi[0].IdParent)
	if err != nil {
		return nil, fmt.Errorf("failed to translate child ID MAP LEGACY HANDLER: %v", err)
	}

	firstNode := Node{
		Child:    rootName,
		Parent:   "",
		Afiliasi: "",
	}

	nodes = append(nodes, firstNode)
	nodeMap[afiliasi.MergedAfiliasi[0].IdParent] = 0

	// Iterate over the AfiliasiList and create/update the nodes
	for _, a := range afiliasi.MergedAfiliasi {
		childName, err := TranslateIdIntoName(db, a.IdChild)
		if err != nil {
			return nil, fmt.Errorf("failed to translate child ID: %v", err)
		}

		parentName, err := TranslateIdIntoName(db, a.IdParent)
		if err != nil {
			return nil, fmt.Errorf("failed to translate parent ID: %v", err)
		}

		if err := updateOrCreateNode(db, a, childName, parentName, a.HubunganAfiliasi, &nodes, &nodeMap); err != nil {
			return nil, err
		}
	}

	return nodes, nil
}

func updateOrCreateNode(db *sql.DB, afiliasi Afiliasi, childName, parentName, hubunganAfiliasi string, nodes *[]Node, nodeMap *map[string]int) error {
	// Create a new node
	node := Node{
		Child:    childName,
		Parent:   parentName,
		Afiliasi: hubunganAfiliasi,
	}

	*nodes = append(*nodes, node)
	(*nodeMap)[afiliasi.IdChild] = len(*nodes) - 1

	// Check for nested afiliasi
	if hasNestedAfiliasi(db, afiliasi.IdChild) {
		if err := processNestedAfiliasi(db, afiliasi.IdChild, nodes, nodeMap); err != nil {
			return fmt.Errorf("failed to process nested afiliasi: %v", err)
		}
	}

	return nil
}

var idToNameMap = make(map[string]string)

func TranslateIdIntoName(db *sql.DB, id string) (string, error) {
	// Check if the translation exists in the cache
	if name, ok := idToNameMap[id]; ok {
		return name, nil
	}

	var name string

	// Check if idParent exists in the data_nasabah table
	err := db.QueryRow("SELECT nama_pengusaha FROM data_nasabah WHERE id = ?", id).Scan(&name)
	if err == nil {
		idToNameMap[id] = name
		return name, nil
	}

	// idParent does not exist in the data_nasabah table, retrieve the name from afiliasi table
	err = db.QueryRow("SELECT nama_child FROM afiliasi WHERE id_child = ?", id).Scan(&name)
	if err == nil {
		idToNameMap[id] = name
		return name, nil
	}

	// idParent does not exist in the afiliasi table, retrieve the name from afiliasi table using id_parent
	err = db.QueryRow("SELECT nama_child FROM afiliasi WHERE id_parent = ?", id).Scan(&name)
	if err != nil {
		log.Printf("Failed to translate child ID '%s': %v", id, err)
		return "", err
	}

	// Cache the translation in the map
	idToNameMap[id] = name

	return name, nil
}

func processNestedAfiliasi(db *sql.DB, idChild string, nodes *[]Node, nodeMap *map[string]int) error {
	// Fetch the nested afiliasi for the given child ID from the database
	nestedAfiliasi, err := fetchNestedAfiliasi(db, idChild)
	if err != nil {
		return fmt.Errorf("failed to fetch nested afiliasi: %v", err)
	}

	// Iterate over the nested afiliasi and create/update the nodes
	for _, nested := range nestedAfiliasi {
		childName := nested.NamaAfiliasi
		parentName, err := TranslateIdIntoName(db, nested.IdParent)
		if err != nil {
			return err
		}

		if err := updateOrCreateNode(db, nested, childName, parentName, nested.HubunganAfiliasi, nodes, nodeMap); err != nil {
			return err
		}
	}

	return nil
}

func hasNestedAfiliasi(db *sql.DB, idChild string) bool {
	// Execute the SQL query to check if there are nested afiliasi for the given child ID
	query := "SELECT COUNT(*) FROM afiliasi WHERE id_parent = ?"
	var count int
	err := db.QueryRow(query, idChild).Scan(&count)
	if err != nil {
		// Handle the error according to your application's error handling approach
		return false
	}

	return count > 0
}

func fetchNestedAfiliasi(db *sql.DB, idChild string) ([]Afiliasi, error) {
	// Execute the SQL query to fetch the nested afiliasi for the given child ID
	query := "SELECT nama_child, id_child, id_parent, hubungan FROM afiliasi WHERE id_parent = ?"
	rows, err := db.Query(query, idChild)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nestedAfiliasi []Afiliasi
	for rows.Next() {
		var afiliasi Afiliasi
		err := rows.Scan(&afiliasi.NamaAfiliasi, &afiliasi.IdChild, &afiliasi.IdParent, &afiliasi.HubunganAfiliasi)
		if err != nil {
			// Handle the error according to your application's error handling approach
			return nil, err
		}
		nestedAfiliasi = append(nestedAfiliasi, afiliasi)
	}

	return nestedAfiliasi, nil
}
