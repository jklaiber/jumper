package inventory

import (
	"fmt"
	"testing"
)

func TestInventoryService_GetHostSSHKey(t *testing.T) {
	inv := DefaultInventory()
	invService := InventoryService{
		Inventory: &inv,
	}

	tests := []struct {
		groupName string
		hostName  string
		want      string
		wantErr   bool
	}{
		{
			groupName: "",
			hostName:  "ungroupedhost1",
			want:      "/home/user/.ssh/id_rsa",
			wantErr:   false,
		},
		{
			groupName: "webservers",
			hostName:  "webserver1",
			want:      "/home/user/.ssh/id_ecdsa",
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver1",
			want:      "/home/user/.ssh/id_rsa",
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver2",
			want:      "/home/user/.ssh/id_rsa",
			wantErr:   false,
		},
		{
			groupName: "",
			hostName:  "nonexistenthost",
			want:      "",
			wantErr:   true,
		},
		{
			groupName: "nonexistentgroup",
			hostName:  "webserver1",
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s in %s", tt.hostName, tt.groupName), func(t *testing.T) {
			got, err := invService.GetHostSSHKey(tt.groupName, tt.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostSSHKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHostSSHKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
