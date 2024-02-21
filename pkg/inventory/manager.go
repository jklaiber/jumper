//go:generate mockgen -source=manager.go -destination=mocks/manager_mock.go -package=mocks InventoryManager
package inventory

import (
	"github.com/jklaiber/jumper/internal/config"
	"github.com/jklaiber/jumper/pkg/access"
)

type HostDetail struct {
	Name string
	Vars Vars
}

type GroupDetail struct {
	Name  string
	Vars  Vars
	Hosts []HostDetail
}

type InventoryManager interface {
	GetUngroupedHosts() []HostDetail
	GetGroups() []GroupDetail
	GetGroupHosts(group string) []HostDetail
	GetHostSSHAgent(groupName, hostName string) (bool, error)
	GetHostSSHAgentForwarding(groupName, hostName string) (bool, error)
	GetHostSSHKey(groupName, hostName string) (string, error)
	GetHostUsername(groupName, hostName string) (string, error)
	GetHostPassword(groupName, hostName string) (string, error)
	GetHostAddress(groupName, hostName string) (string, error)
	GetHostPort(groupName, hostName string) (int, error)
	GetAccessConfig(groupName, hostName string) (*access.AccessConfig, error)
}

type InventoryService struct {
	Inventory *Inventory
}

func NewInventoryService(reader InventoryReader, parser InventoryParser) (*InventoryService, error) {
	invStr, err := reader.ReadInventory(config.Params.InventoryPath)
	if err != nil {
		return nil, err
	}
	inv, err := parser.Parse([]byte(invStr))
	if err != nil {
		return nil, err
	}

	return &InventoryService{
		Inventory: inv,
	}, nil
}
