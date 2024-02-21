package inventory

func (s *InventoryService) GetUngroupedHosts() []HostDetail {
	var ungroupedHosts []HostDetail
	for key, vars := range s.Inventory.All.Hosts {
		ungroupedHosts = append(ungroupedHosts, HostDetail{Name: key, Vars: vars})
	}
	return ungroupedHosts
}

func (s *InventoryService) GetGroups() []GroupDetail {
	var groups []GroupDetail
	for key, children := range s.Inventory.All.Children {
		var hosts []HostDetail
		for host, vars := range children.Hosts {
			hosts = append(hosts, HostDetail{Name: host, Vars: vars})
		}
		groups = append(groups, GroupDetail{Name: key, Vars: children.Vars, Hosts: hosts})
	}
	return groups
}

func (s *InventoryService) GetGroupHosts(group string) []HostDetail {
	var hosts []HostDetail
	if groupDetail, exists := s.Inventory.All.Children[group]; exists {
		for host, vars := range groupDetail.Hosts {
			hosts = append(hosts, HostDetail{Name: host, Vars: vars})
		}
	}
	return hosts
}
