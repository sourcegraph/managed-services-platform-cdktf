package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterSecurityConfig struct {
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_admin_cluster#authorization GkeonpremBareMetalAdminCluster#authorization}
	Authorization *GkeonpremBareMetalAdminClusterSecurityConfigAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
}

