package zerotrustaccesspolicy


type ZeroTrustAccessPolicyIncludeDevicePosture struct {
	// The ID of a device posture integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_policy#integration_uid ZeroTrustAccessPolicy#integration_uid}
	IntegrationUid *string `field:"required" json:"integrationUid" yaml:"integrationUid"`
}

