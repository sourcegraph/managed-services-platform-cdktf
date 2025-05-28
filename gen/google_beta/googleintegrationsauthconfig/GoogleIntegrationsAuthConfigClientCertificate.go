package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigClientCertificate struct {
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_integrations_auth_config#encrypted_private_key GoogleIntegrationsAuthConfig#encrypted_private_key}
	EncryptedPrivateKey *string `field:"required" json:"encryptedPrivateKey" yaml:"encryptedPrivateKey"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_integrations_auth_config#ssl_certificate GoogleIntegrationsAuthConfig#ssl_certificate}
	SslCertificate *string `field:"required" json:"sslCertificate" yaml:"sslCertificate"`
	// 'passphrase' should be left unset if private key is not encrypted.
	//
	// Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_integrations_auth_config#passphrase GoogleIntegrationsAuthConfig#passphrase}
	Passphrase *string `field:"optional" json:"passphrase" yaml:"passphrase"`
}

