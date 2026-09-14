package googlecontainercluster


type GoogleContainerClusterAddonsConfigLustreCsiDriverConfig struct {
	// Whether the Lustre CSI driver is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// If set to true, the Lustre CSI driver will initialize LNet (the virtual network layer for Lustre kernel module) using port 6988.
	//
	// This flag is required to workaround a port conflict with the gke-metadata-server on GKE nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_container_cluster#enable_legacy_lustre_port GoogleContainerCluster#enable_legacy_lustre_port}
	EnableLegacyLustrePort interface{} `field:"optional" json:"enableLegacyLustrePort" yaml:"enableLegacyLustrePort"`
}

