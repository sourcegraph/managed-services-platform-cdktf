package identityplatformconfig


type IdentityPlatformConfigClient struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/identity_platform_config#permissions IdentityPlatformConfig#permissions}
	Permissions *IdentityPlatformConfigClientPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

