package googleprivatecacapool

type GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage struct {
	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#client_auth GooglePrivatecaCaPool#client_auth}
	ClientAuth any `field:"optional" json:"clientAuth" yaml:"clientAuth"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#code_signing GooglePrivatecaCaPool#code_signing}
	CodeSigning any `field:"optional" json:"codeSigning" yaml:"codeSigning"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#email_protection GooglePrivatecaCaPool#email_protection}
	EmailProtection any `field:"optional" json:"emailProtection" yaml:"emailProtection"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#ocsp_signing GooglePrivatecaCaPool#ocsp_signing}
	OcspSigning any `field:"optional" json:"ocspSigning" yaml:"ocspSigning"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#server_auth GooglePrivatecaCaPool#server_auth}
	ServerAuth any `field:"optional" json:"serverAuth" yaml:"serverAuth"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#time_stamping GooglePrivatecaCaPool#time_stamping}
	TimeStamping any `field:"optional" json:"timeStamping" yaml:"timeStamping"`
}
