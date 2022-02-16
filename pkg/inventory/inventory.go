package inventory

import (
	"errors"
	"io/ioutil"

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
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SshKey   string `yaml:"sshkey"`
	Ip       string `yaml:"ip"`
}

func NewInventory(filePath string) (Inventory, error) {
	inventory := Inventory{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return inventory, errors.New("inventory file could not be read")
	}
	err = yaml.Unmarshal([]byte(b), &inventory)
	if err != nil {
		return inventory, errors.New("inventory could not be unmarshalled")
	}
	return inventory, nil
}

func (inventory *Inventory) GetAccessInformation(group string, host string) (username string, password string, sshkey string, ip string, err error) {
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
	ip, err = inventory.GetIp(group, host)
	if err != nil {
		return "", "", "", "", errors.New("ip for host not found")
	}
	return
}
