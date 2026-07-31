package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesExcludeDevicePosture struct {
	// The ID of a device posture integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_application#integration_uid ZeroTrustAccessApplication#integration_uid}
	IntegrationUid *string `field:"required" json:"integrationUid" yaml:"integrationUid"`
}

