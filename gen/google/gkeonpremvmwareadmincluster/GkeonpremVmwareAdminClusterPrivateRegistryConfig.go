package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterPrivateRegistryConfig struct {
	// The registry address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_admin_cluster#address GkeonpremVmwareAdminCluster#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The CA certificate public key for private registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_admin_cluster#ca_cert GkeonpremVmwareAdminCluster#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
}

