package googleidentityplatformtenant


type GoogleIdentityPlatformTenantClientPermissions struct {
	// When true, end users cannot delete their account on the associated project through any of our API methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_identity_platform_tenant#disabled_user_deletion GoogleIdentityPlatformTenant#disabled_user_deletion}
	DisabledUserDeletion interface{} `field:"optional" json:"disabledUserDeletion" yaml:"disabledUserDeletion"`
	// When true, end users cannot sign up for a new account on the associated project through any of our API methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_identity_platform_tenant#disabled_user_signup GoogleIdentityPlatformTenant#disabled_user_signup}
	DisabledUserSignup interface{} `field:"optional" json:"disabledUserSignup" yaml:"disabledUserSignup"`
}

