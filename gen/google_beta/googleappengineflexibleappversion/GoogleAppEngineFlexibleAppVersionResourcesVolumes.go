package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionResourcesVolumes struct {
	// Unique name for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_flexible_app_version#name GoogleAppEngineFlexibleAppVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Volume size in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_flexible_app_version#size_gb GoogleAppEngineFlexibleAppVersion#size_gb}
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Underlying volume type, e.g. 'tmpfs'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_flexible_app_version#volume_type GoogleAppEngineFlexibleAppVersion#volume_type}
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
}

