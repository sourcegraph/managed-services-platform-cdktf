package googleidentityplatformconfig


type GoogleIdentityPlatformConfigClient struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_identity_platform_config#permissions GoogleIdentityPlatformConfig#permissions}
	Permissions *GoogleIdentityPlatformConfigClientPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

