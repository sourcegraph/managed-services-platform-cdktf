package googleidentityplatformtenant


type GoogleIdentityPlatformTenantClient struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_identity_platform_tenant#permissions GoogleIdentityPlatformTenant#permissions}
	Permissions *GoogleIdentityPlatformTenantClientPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

