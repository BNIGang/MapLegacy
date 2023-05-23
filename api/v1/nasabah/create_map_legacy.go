package v1

import (
	"fmt"

	"github.com/BNIGang/MapLegacy/web"
)

type Node struct {
	Child  string `json:"child"`
	Parent string `json:"parent"`
	Spouse string `json:"spouse,omitempty"`
}

func MapLegacyHandler(afiliasi *MergedRow) ([]Node, error) {
	// Check if afiliasi is empty
	if len(afiliasi.MergedAfiliasi) == 0 {
		// Create a single empty node
		node := Node{
			Child:  "",
			Parent: "",
		}
		return []Node{node}, nil
	}

	var nodes []Node
	nodeMap := make(map[string]int) // Map to store the index of existing nodes by child ID

	// Add the first node with Child as a.IdParent and Parent as an empty string
	childName, err := TranslateIdIntoName(afiliasi.MergedAfiliasi[0].IdParent)
	if err != nil {
		return nil, fmt.Errorf("failed to translate child ID: %v", err)
	}

	firstNode := Node{
		Child:  childName,
		Parent: "",
	}
	nodes = append(nodes, firstNode)
	nodeMap[afiliasi.MergedAfiliasi[0].IdParent] = 0

	// Iterate over the remaining AfiliasiList and create/update the nodes
	for _, a := range afiliasi.MergedAfiliasi {
		childName, err := TranslateIdIntoName(a.IdChild)
		if err != nil {
			return nil, fmt.Errorf("failed to translate child ID: %v", err)
		}

		parentName, err := TranslateIdIntoName(a.IdParent)
		if err != nil {
			return nil, fmt.Errorf("failed to translate parent ID: %v", err)
		}

		// Check if the HubunganAfiliasi is "Istri" or "Suami"
		if a.HubunganAfiliasi == "Istri" || a.HubunganAfiliasi == "Suami" {
			// Find the index of the existing node with the matching child ID
			idx, exists := nodeMap[a.IdParent]
			if exists {
				// Update the existing node's spouse field with the current parent name
				nodes[idx].Spouse = childName
				continue // Skip creating a new node
			}
		}

		node := Node{
			Child:  childName,
			Parent: parentName,
		}

		nodes = append(nodes, node)
		nodeMap[a.IdChild] = len(nodes) - 1
	}

	return nodes, nil
}

var idToNameMap = make(map[string]string)

func TranslateIdIntoName(id string) (string, error) {
	// Check if the translation exists in the cache
	if name, ok := idToNameMap[id]; ok {
		return name, nil
	}

	var name string

	db, err := web.Connect()
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Check if idParent exists in the data_nasabah table
	err = db.QueryRow("SELECT nama_pengusaha FROM data_nasabah WHERE id = ?", id).Scan(&name)
	if err == nil {
		return name, nil
	}

	// idParent does not exist in the data_nasabah table, retrieve the name from afiliasi table
	err = db.QueryRow("SELECT nama_child FROM afiliasi WHERE id_child = ?", id).Scan(&name)
	if err != nil {
		return "", err
	}

	// Cache the translation in the map
	idToNameMap[id] = name

	return name, nil
}
