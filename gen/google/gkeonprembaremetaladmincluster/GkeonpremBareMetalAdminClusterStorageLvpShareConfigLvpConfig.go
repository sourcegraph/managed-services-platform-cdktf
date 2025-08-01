package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfig struct {
	// The host machine path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_admin_cluster#path GkeonpremBareMetalAdminCluster#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The StorageClass name that PVs will be created with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_admin_cluster#storage_class GkeonpremBareMetalAdminCluster#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
}

