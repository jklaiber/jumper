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
	Username       string `yaml:"username,omitempty"`
	AnsibleUser    string `yaml:"ansible_user,omitempty"`
	Password       string `yaml:"password,omitempty"`
	AnsibleSshPASS string `yaml:"ansible_ssh_pass,omitempty"`
	SshKey         string `yaml:"sshkey,omitempty"`
	Address        string `yaml:"address,omitempty"`
	AnsibleHost    string `yaml:"ansible_host,omitempty"`
	SshAgent       bool   `yaml:"sshagent,omitempty"`
	Port           int    `yaml:"port,omitempty"`
}

func NewInventory(filePath string) (inventory Inventory, err error) {
	inventory = Inventory{}
	// str, err := ioutil.ReadFile(filePath)
	str, err := inventory.readEncryptedFile(filePath)
	if err != nil {
		return inventory, errors.New("inventory file could not be read")
	}
	err = yaml.Unmarshal([]byte(str), &inventory)
	if err != nil {
		return inventory, errors.New("inventory could not be unmarshalled")
	}
	return
}

func (inventory *Inventory) GetAccessInformation(group string, host string) (username string, password string, sshkey string, sshagent bool, address string, port int, err error) {
	username, err = inventory.GetUsername(group, host)
	if err != nil {
		return "", "", "", false, "", 0, errors.New("username for host not found")
	}
	password, sshkey, sshagent, err = inventory.getAccessMethod(group, host)
	if err != nil {
		return "", "", "", false, "", 0, errors.New("no valid access method found")
	}
	address, port, err = inventory.GetAddress(group, host)
	if err != nil {
		return "", "", "", false, "", 0, errors.New("no valid address found")
	}
	return username, password, sshkey, sshagent, address, port, nil
}

func (inventory *Inventory) getAccessMethod(group string, host string) (string, string, bool, error) {
	password, passerr := inventory.GetPassword(group, host)
	sshkey, sshkeyerr := inventory.GetSshKey(group, host)
	sshagent, _ := inventory.GetSshAgent(group, host)
	if passerr == nil || sshkeyerr == nil || sshagent {
		return password, sshkey, sshagent, nil
	}
	return "", "", false, errors.New("no valid access method found")
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
