package identityplatformconfig


type IdentityPlatformConfigSmsRegionConfig struct {
	// allow_by_default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/identity_platform_config#allow_by_default IdentityPlatformConfig#allow_by_default}
	AllowByDefault *IdentityPlatformConfigSmsRegionConfigAllowByDefault `field:"optional" json:"allowByDefault" yaml:"allowByDefault"`
	// allowlist_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/identity_platform_config#allowlist_only IdentityPlatformConfig#allowlist_only}
	AllowlistOnly *IdentityPlatformConfigSmsRegionConfigAllowlistOnly `field:"optional" json:"allowlistOnly" yaml:"allowlistOnly"`
}

