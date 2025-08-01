package zerotrustaccessapplication


type ZeroTrustAccessApplicationScimConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#idp_uid ZeroTrustAccessApplication#idp_uid}
	IdpUid *string `field:"required" json:"idpUid" yaml:"idpUid"`
	// The base URI for the application's SCIM-compatible API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#remote_uri ZeroTrustAccessApplication#remote_uri}
	RemoteUri *string `field:"required" json:"remoteUri" yaml:"remoteUri"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#authentication ZeroTrustAccessApplication#authentication}
	Authentication *ZeroTrustAccessApplicationScimConfigAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM resources.
	//
	// If true, sets 'active' to false on the SCIM resource. Note: Some targets do not support DELETE operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#deactivate_on_delete ZeroTrustAccessApplication#deactivate_on_delete}
	DeactivateOnDelete interface{} `field:"optional" json:"deactivateOnDelete" yaml:"deactivateOnDelete"`
	// Whether SCIM provisioning is turned on for this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#enabled ZeroTrustAccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this application.
	//
	// These can transform or filter the resources to be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#mappings ZeroTrustAccessApplication#mappings}
	Mappings interface{} `field:"optional" json:"mappings" yaml:"mappings"`
}

