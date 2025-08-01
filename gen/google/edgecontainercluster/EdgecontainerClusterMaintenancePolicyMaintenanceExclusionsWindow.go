package edgecontainercluster


type EdgecontainerClusterMaintenancePolicyMaintenanceExclusionsWindow struct {
	// The time that the window ends. The end time must take place after the start time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#end_time EdgecontainerCluster#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time that the window first starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#start_time EdgecontainerCluster#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

