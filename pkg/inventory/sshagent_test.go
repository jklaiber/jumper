package inventory

import (
	"fmt"
	"testing"
)

func TestInventoryService_GetHostSSHAgent(t *testing.T) {
	inv := DefaultInventory()
	invService := InventoryService{
		Inventory: &inv,
	}

	tests := []struct {
		groupName string
		hostName  string
		want      bool
		wantErr   bool
	}{
		{
			groupName: "",
			hostName:  "ungroupedhost1",
			want:      false,
			wantErr:   false,
		},
		{
			groupName: "webservers",
			hostName:  "webserver1",
			want:      false,
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver1",
			want:      false,
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver2",
			want:      true,
			wantErr:   false,
		},
		{
			groupName: "",
			hostName:  "nonexistenthost",
			want:      false,
			wantErr:   true,
		},
		{
			groupName: "nonexistentgroup",
			hostName:  "webserver1",
			want:      false,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s in %s", tt.hostName, tt.groupName), func(t *testing.T) {
			got, err := invService.GetHostSSHAgent(tt.groupName, tt.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostSSHAgent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHostSSHAgent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryService_GetHostSSHAgentForwarding(t *testing.T) {
	inv := DefaultInventory()
	invService := InventoryService{
		Inventory: &inv,
	}

	tests := []struct {
		groupName string
		hostName  string
		want      bool
		wantErr   bool
	}{
		{
			groupName: "",
			hostName:  "ungroupedhost1",
			want:      false,
			wantErr:   false,
		},
		{
			groupName: "webservers",
			hostName:  "webserver1",
			want:      false,
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver1",
			want:      false,
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver2",
			want:      true,
			wantErr:   false,
		},
		{
			groupName: "",
			hostName:  "nonexistenthost",
			want:      false,
			wantErr:   true,
		},
		{
			groupName: "nonexistentgroup",
			hostName:  "webserver1",
			want:      false,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s in %s", tt.hostName, tt.groupName), func(t *testing.T) {
			got, err := invService.GetHostSSHAgentForwarding(tt.groupName, tt.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostSSHAgentForwarding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHostSSHAgentForwarding() got = %v, want %v", got, tt.want)
			}
		})
	}
}
