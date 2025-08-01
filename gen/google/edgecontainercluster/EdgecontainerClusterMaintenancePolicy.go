package edgecontainercluster


type EdgecontainerClusterMaintenancePolicy struct {
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#window EdgecontainerCluster#window}
	Window *EdgecontainerClusterMaintenancePolicyWindow `field:"required" json:"window" yaml:"window"`
	// maintenance_exclusions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#maintenance_exclusions EdgecontainerCluster#maintenance_exclusions}
	MaintenanceExclusions interface{} `field:"optional" json:"maintenanceExclusions" yaml:"maintenanceExclusions"`
}

