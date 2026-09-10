package zerotrustorganization


type ZeroTrustOrganizationMfaSshPivKeyRequirements struct {
	// Defines when a PIN is required to use the SSH key.
	//
	// Valid values: `never` (no PIN required), `once` (PIN required once per session), `always` (PIN required for each use).
	// Available values: "never", "once", "always".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#pin_policy ZeroTrustOrganization#pin_policy}
	PinPolicy *string `field:"optional" json:"pinPolicy" yaml:"pinPolicy"`
	// Requires the SSH PIV key to be stored on a FIPS 140-2 Level 1 or higher validated device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#require_fips_device ZeroTrustOrganization#require_fips_device}
	RequireFipsDevice interface{} `field:"optional" json:"requireFipsDevice" yaml:"requireFipsDevice"`
	// Specifies the allowed SSH key sizes in bits.
	//
	// Valid sizes depend on key type. Ed25519 has a fixed key size and does not accept this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#ssh_key_size ZeroTrustOrganization#ssh_key_size}
	SshKeySize *[]*float64 `field:"optional" json:"sshKeySize" yaml:"sshKeySize"`
	// Specifies the allowed SSH key types. Valid values are `ecdsa`, `ed25519`, and `rsa`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#ssh_key_type ZeroTrustOrganization#ssh_key_type}
	SshKeyType *[]*string `field:"optional" json:"sshKeyType" yaml:"sshKeyType"`
	// Defines when physical touch is required to use the SSH key.
	//
	// Valid values: `never` (no touch required), `always` (touch required for each use), `cached` (touch cached for 15 seconds).
	// Available values: "never", "always", "cached".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#touch_policy ZeroTrustOrganization#touch_policy}
	TouchPolicy *string `field:"optional" json:"touchPolicy" yaml:"touchPolicy"`
}

