package zerotrustaccesspolicy


type ZeroTrustAccessPolicyRequireLinkedAppToken struct {
	// The ID of an Access OIDC SaaS application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_policy#app_uid ZeroTrustAccessPolicy#app_uid}
	AppUid *string `field:"required" json:"appUid" yaml:"appUid"`
}

