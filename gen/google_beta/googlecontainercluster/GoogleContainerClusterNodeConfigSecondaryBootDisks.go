package googlecontainercluster


type GoogleContainerClusterNodeConfigSecondaryBootDisks struct {
	// Disk image to create the secondary boot disk from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#disk_image GoogleContainerCluster#disk_image}
	DiskImage *string `field:"required" json:"diskImage" yaml:"diskImage"`
	// Mode for how the secondary boot disk is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#mode GoogleContainerCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

