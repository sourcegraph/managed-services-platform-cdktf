package googlestoragecontrolfolderintelligenceconfig


type GoogleStorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations struct {
	// List of locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_control_folder_intelligence_config#locations GoogleStorageControlFolderIntelligenceConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

