package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterMaintenanceConfig struct {
	// All IPv4 address from these ranges will be placed into maintenance mode.
	//
	// Nodes in maintenance mode will be cordoned and drained. When both of these
	// are true, the "baremetal.cluster.gke.io/maintenance" annotation will be set
	// on the node resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#maintenance_address_cidr_blocks GoogleGkeonpremBareMetalAdminCluster#maintenance_address_cidr_blocks}
	MaintenanceAddressCidrBlocks *[]*string `field:"required" json:"maintenanceAddressCidrBlocks" yaml:"maintenanceAddressCidrBlocks"`
}

