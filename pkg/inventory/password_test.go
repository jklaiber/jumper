package inventory

import (
	"fmt"
	"testing"
)

func TestInventoryService_GetHostPassword(t *testing.T) {
	inv := DefaultInventory()
	invService := InventoryService{
		Inventory: &inv,
	}

	tests := []struct {
		groupName    string
		hostName     string
		wantPassword string
		wantErr      bool
	}{
		{
			groupName:    "",
			hostName:     "ungroupedhost1",
			wantPassword: "ungroupedpassword",
			wantErr:      false,
		},
		{
			groupName:    "webservers",
			hostName:     "webserver1",
			wantPassword: "bar",
			wantErr:      false,
		},
		{
			groupName:    "dbservers",
			hostName:     "dbserver1",
			wantPassword: "bar",
			wantErr:      false,
		},
		{
			groupName:    "",
			hostName:     "nonexistenthost",
			wantPassword: "",
			wantErr:      true,
		},
		{
			groupName:    "nonexistentgroup",
			hostName:     "webserver1",
			wantPassword: "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s in %s", tt.hostName, tt.groupName), func(t *testing.T) {
			gotPassword, err := invService.GetHostPassword(tt.groupName, tt.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPassword != tt.wantPassword {
				t.Errorf("GetHostPassword() gotPassword = %v, want %v", gotPassword, tt.wantPassword)
			}
		})
	}
}
