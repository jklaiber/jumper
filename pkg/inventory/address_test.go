package inventory

import (
	"fmt"
	"testing"
)

func TestInventoryService_GetHostAddress(t *testing.T) {
	inv := DefaultInventory()
	invService := InventoryService{
		Inventory: &inv,
	}

	tests := []struct {
		groupName string
		hostName  string
		wantAddr  string
		wantErr   bool
	}{
		{
			groupName: "",
			hostName:  "ungroupedhost1",
			wantAddr:  "ungroupedhost1.example.com",
			wantErr:   false,
		},
		{
			groupName: "webservers",
			hostName:  "webserver1",
			wantAddr:  "webserver1.example.com",
			wantErr:   false,
		},
		{
			groupName: "dbservers",
			hostName:  "dbserver1",
			wantAddr:  "192.168.1.10",
			wantErr:   false,
		},
		{
			groupName: "",
			hostName:  "nonexistenthost",
			wantAddr:  "",
			wantErr:   true,
		},
		{
			groupName: "nonexistentgroup",
			hostName:  "webserver1",
			wantAddr:  "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s in %s", tt.hostName, tt.groupName), func(t *testing.T) {
			gotAddr, err := invService.GetHostAddress(tt.groupName, tt.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAddr != tt.wantAddr {
				t.Errorf("GetHostAddress() gotAddr = %v, want %v", gotAddr, tt.wantAddr)
			}
		})
	}
}
