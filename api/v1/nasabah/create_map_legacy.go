package v1

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Child  string `json:"child"`
	Parent string `json:"parent"`
}

func MapLegacyHandler(afiliasi *MergedRow) ([]byte, error) {
	// Check if afiliasi is empty
	if len(afiliasi.MergedAfiliasi) == 0 {
		// Create a single empty node
		node := Node{
			Child:  "",
			Parent: "",
		}
		// Convert the empty node to JSON
		data, err := json.Marshal([]Node{node})
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %v", err)
		}
		return data, nil
	}

	var nodes []Node

	// Iterate over the AfiliasiList and create the nodes
	for _, a := range afiliasi.MergedAfiliasi {
		node := Node{
			Child:  a.IdChild,
			Parent: a.IdParent,
		}
		nodes = append(nodes, node)
	}

	// Convert nodes to JSON
	data, err := json.Marshal(nodes)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return data, nil
}
