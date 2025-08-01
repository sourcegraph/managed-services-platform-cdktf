package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterStorageLvpShareConfig struct {
	// lvp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_admin_cluster#lvp_config GkeonpremBareMetalAdminCluster#lvp_config}
	LvpConfig *GkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfig `field:"required" json:"lvpConfig" yaml:"lvpConfig"`
	// The number of subdirectories to create under path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_admin_cluster#shared_path_pv_count GkeonpremBareMetalAdminCluster#shared_path_pv_count}
	SharedPathPvCount *float64 `field:"optional" json:"sharedPathPvCount" yaml:"sharedPathPvCount"`
}

