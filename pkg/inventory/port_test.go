package inventory

import (
	"fmt"
	"testing"
)

func TestInventoryService_GetHostPort(t *testing.T) {
	inv := DefaultInventory()
	invService := InventoryService{
		Inventory: &inv,
	}

	tests := []struct {
		groupName string
		hostName  string
		wantPort  int
		wantErr   bool
	}{
		{
			groupName: "webservers",
			hostName:  "webserver1",
			wantPort:  22,
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver1",
			wantPort:  22,
			wantErr:   false,
		},
		{
			groupName: "webservers",
			hostName:  "webserver2",
			wantPort:  2222,
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver2",
			wantPort:  22,
			wantErr:   false,
		},
		{
			groupName: "",
			hostName:  "ungroupedhost1",
			wantPort:  22,
			wantErr:   false,
		},
		{
			groupName: "",
			hostName:  "nonexistenthost",
			wantPort:  0,
			wantErr:   true,
		},
		{
			groupName: "nonexistentgroup",
			hostName:  "webserver1",
			wantPort:  0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s in %s", tt.hostName, tt.groupName), func(t *testing.T) {
			gotPort, err := invService.GetHostPort(tt.groupName, tt.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPort != tt.wantPort {
				t.Errorf("GetHostPort() gotPort = %v, want %v", gotPort, tt.wantPort)
			}
		})
	}
}
