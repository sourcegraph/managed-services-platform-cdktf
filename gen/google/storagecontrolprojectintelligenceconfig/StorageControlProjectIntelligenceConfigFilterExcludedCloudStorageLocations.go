package storagecontrolprojectintelligenceconfig


type StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations struct {
	// List of locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_control_project_intelligence_config#locations StorageControlProjectIntelligenceConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

