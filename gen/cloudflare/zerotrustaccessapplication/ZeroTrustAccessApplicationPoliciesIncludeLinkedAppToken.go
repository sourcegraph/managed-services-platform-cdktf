package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesIncludeLinkedAppToken struct {
	// The ID of an Access OIDC SaaS application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_application#app_uid ZeroTrustAccessApplication#app_uid}
	AppUid *string `field:"required" json:"appUid" yaml:"appUid"`
}

