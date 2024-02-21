package inventory

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
	Username           string `yaml:"username,omitempty"`
	AnsibleUser        string `yaml:"ansible_user,omitempty"`
	Password           string `yaml:"password,omitempty"`
	AnsibleSshPASS     string `yaml:"ansible_ssh_pass,omitempty"`
	SshKey             string `yaml:"sshkey,omitempty"`
	Address            string `yaml:"address,omitempty"`
	AnsibleHost        string `yaml:"ansible_host,omitempty"`
	SshAgent           bool   `yaml:"sshagent,omitempty"`
	SshAgentForwarding bool   `yaml:"sshagent_forwarding,omitempty"`
	Port               int    `yaml:"port,omitempty"`
}

func DefaultInventory() Inventory {
	return Inventory{
		All: HostInventory{
			Hosts: map[string]Vars{
				"ungroupedhost1": {},
				"ungroupedhost2": {},
			},
			Children: map[string]ChildrenGroup{
				"webservers": {
					Hosts: map[string]Vars{
						"webserver1": {Address: "webserver1.example.com"},
						"webserver2": {Address: "webserver2.example.com"},
					},
					Vars: Vars{Username: "foo", Password: "bar"},
				},
				"dbservers": {
					Hosts: map[string]Vars{
						"dbserver1": {Address: "192.168.1.10", Username: "foo", Password: "bar"},
						"dbserver2": {Address: "192.168.1.11", Username: "foo", SshAgent: true, SshAgentForwarding: true},
					},
				},
				"fileserver": {
					Hosts: map[string]Vars{
						"fileserver1": {},
					},
				},
			},
			Vars: Vars{SshKey: "/home/user/.ssh/id_rsa", Username: "globalusername"},
		},
	}
}
