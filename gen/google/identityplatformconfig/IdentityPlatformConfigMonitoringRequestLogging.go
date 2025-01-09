package identityplatformconfig


type IdentityPlatformConfigMonitoringRequestLogging struct {
	// Whether logging is enabled for this project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/identity_platform_config#enabled IdentityPlatformConfig#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

