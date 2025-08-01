package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterStorageLvpShareConfigLvpConfig struct {
	// The host machine path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#path GkeonpremBareMetalCluster#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The StorageClass name that PVs will be created with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#storage_class GkeonpremBareMetalCluster#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
}

