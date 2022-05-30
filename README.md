![jumper logo](doc/jumper_logo.png)

Jumper is a simple and fast cli ssh manager. The goal was to reuse some ideas like the Ansible YAML inventory and Ansible Vault.

## Installation

You can use the following command to install Jumper to your system:
```
sh -c "$(curl -fsSL https://raw.githubusercontent.com/jklaiber/jumper/main/install/install.sh)"
```

## Usage
The tool can be used to connect over ssh to the saved hosts in the inventory.   
  
Connect for example to an ungrouped host:
```
$ jumper connect ungroupedhost1
```
Or to a host which belongs to a group:
```
$ jumper connect -g webservers webserver1
```

## Configuration

### Autocompletion
In order to use the autocompletion feature you have to generate the autocompletion file for your specific shell.  
  
Example for bash:
```
$ jumper completion bash >> jumper_completion
```
You can then source the completion file into your actual session (only available until you close the session) or reference it, for example in your `.bashrc` file. 
```
$ source jumper_completion
```

## Inventory
The inventory is completely inspired by the Yaml Ansible inventory (more [here]()).  
  
The basic structure is as follows:
```yaml
all:
  hosts:
    ungroupedhost1:
    ungroupedhost2:
  children:
    webservers:
      hosts:
        webserver1:
          address: webserver1.example.com
        webserver2:
          address: webserver2.example.com
      vars:
        username: foo
        password: bar
    dbservers:
      hosts:
        dbserver1:
          address: 192.168.1.10
          username: foo
          password: bar
        dbserver2:
          address: 192.168.1.11
          username: foo
          sshagent: True
    fileserver:
      hosts:
        fileserver1:
  vars:
    sshkey: /home/user/.ssh/id_rsa
    username: globalusername
```

### Variables Inheritance  
More generic variables will be automatically applied to the hosts. The prioritization rule is as follows:
1. Direct host variables
2. Group variables
3. Global variables
  
### Authentication Methods
The are three different authentication methods implemented which can be used:
* `password` authentication (password as `string`)
* `ssh key` authentication (path to ssh key)
* `ssh agent` authentication (`bool`)

The prioritization rule is as follows:
1. `ssh agent` 
2. `ssh key` 
3. `password` 

In the example below the choosen authentication method would be `sshagent`.

```yaml
all:
  children:
    webservers:
      hosts:
        foo.example.com:
          username: userfoo
          password: passfoo
          sshagent: True
  vars:
    sshkey: /home/user/.ssh/id_rsa
    username: globalusername
```

### Encryption
The whole encryption is done with Ansible Vault (more [here]()).  
  
Through the usage of the same mechanism like Ansible Vault, there is also the possiblity to use `ansible-vault` to edit the inventory file:
```
$ ansible-vault edit inventory.yaml
```
The inventory is completely encrypted and can therefore also be synchronised to a e.g. cloud file service like Onedrive or GoogleDrive.