package inventory

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type InventoryParser interface {
	Parse(data []byte) (*Inventory, error)
}

type DefaultInventoryParser struct{}

func (parser *DefaultInventoryParser) Parse(data []byte) (*Inventory, error) {
	var inv Inventory
	err := yaml.Unmarshal(data, &inv)
	if err != nil {
		return nil, fmt.Errorf("could not parse inventory: %v", err)
	}
	return &inv, nil
}
