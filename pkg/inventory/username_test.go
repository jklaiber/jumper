package inventory

import (
	"fmt"
	"testing"
)

func TestInventoryService_GetHostUsername(t *testing.T) {
	inv := DefaultInventory()
	invService := InventoryService{
		Inventory: &inv,
	}

	tests := []struct {
		groupName    string
		hostName     string
		wantUsername string
		wantErr      bool
	}{
		{
			groupName:    "",
			hostName:     "ungroupedhost1",
			wantUsername: "ungroupedusername",
			wantErr:      false,
		},
		{
			groupName:    "webservers",
			hostName:     "webserver1",
			wantUsername: "foo",
			wantErr:      false,
		},
		{
			groupName:    "dbservers",
			hostName:     "dbserver1",
			wantUsername: "foo",
			wantErr:      false,
		},
		{
			groupName:    "",
			hostName:     "nonexistenthost",
			wantUsername: "",
			wantErr:      true,
		},
		{
			groupName:    "nonexistentgroup",
			hostName:     "webserver1",
			wantUsername: "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s in %s", tt.hostName, tt.groupName), func(t *testing.T) {
			gotUsername, err := invService.GetHostUsername(tt.groupName, tt.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUsername != tt.wantUsername {
				t.Errorf("GetHostUsername() gotUsername = %v, want %v", gotUsername, tt.wantUsername)
			}
		})
	}
}
