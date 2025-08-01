package identityplatformconfig


type IdentityPlatformConfigSignInAnonymous struct {
	// Whether anonymous user auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/identity_platform_config#enabled IdentityPlatformConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

