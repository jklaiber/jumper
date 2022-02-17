package inventory

import (
	"errors"

	"gopkg.in/yaml.v2"
)

type Inventory struct {
	All HostInventory `yaml:"all"`
}

type HostInventory struct {
	Hosts    map[string]Vars          `yaml:"hosts"`
	Children map[string]ChildrenGroup `yaml:"children"`
	Vars     Vars                     `yaml:"vars"`
}

type ChildrenGroup struct {
	Hosts map[string]Vars `yaml:"hosts"`
	Vars  Vars            `yaml:"vars"`
}

type Vars struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	SshKey   string `yaml:"sshkey,omitempty"`
	Address  string `yaml:"address,omitempty"`
}

func NewInventory(filePath string, password string) (inventory Inventory, err error) {
	inventory = Inventory{}
	// str, err := ioutil.ReadFile(filePath)
	str, err := inventory.readEncryptedFile(filePath, password)
	if err != nil {
		return inventory, errors.New("inventory file could not be read")
	}
	err = yaml.Unmarshal([]byte(str), &inventory)
	if err != nil {
		return inventory, errors.New("inventory could not be unmarshalled")
	}
	return
}

func (inventory *Inventory) GetAccessInformation(group string, host string) (username string, password string, sshkey string, address string, err error) {
	username, err = inventory.GetUsername(group, host)
	if err != nil {
		return "", "", "", "", errors.New("username for host not found")
	}
	password, err = inventory.GetPassword(group, host)
	if err != nil {
		return "", "", "", "", errors.New("password for host not found")
	}
	sshkey, err = inventory.GetSshKey(group, host)
	if err != nil {
		return "", "", "", "", errors.New("sshkey for host not found")
	}
	address, err = inventory.GetAddress(group, host)
	if err != nil {
		return "", "", "", "", errors.New("address for host not found")
	}
	return
}

func (inventory *Inventory) GetUngroupedHosts() (ungroupedHosts []string) {
	for key, _ := range inventory.All.Hosts {
		ungroupedHosts = append(ungroupedHosts, key)
	}
	return
}

func (inventory *Inventory) GetGroups() (groups []string) {
	for key, _ := range inventory.All.Children {
		groups = append(groups, key)
	}
	return
}

func (inventory *Inventory) GetGroupHosts(group string) (groupHosts []string) {
	for key, _ := range inventory.All.Children[group].Hosts {
		groupHosts = append(groupHosts, key)
	}
	return
}
