package containercluster


type ContainerClusterNodeConfigHostMaintenancePolicy struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#maintenance_interval ContainerCluster#maintenance_interval}
	MaintenanceInterval *string `field:"required" json:"maintenanceInterval" yaml:"maintenanceInterval"`
}

