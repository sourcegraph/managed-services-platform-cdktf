package identityplatformtenant


type IdentityPlatformTenantClient struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/identity_platform_tenant#permissions IdentityPlatformTenant#permissions}
	Permissions *IdentityPlatformTenantClientPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

