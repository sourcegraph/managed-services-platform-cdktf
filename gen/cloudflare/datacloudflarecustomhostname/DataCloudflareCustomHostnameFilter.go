package datacloudflarecustomhostname


type DataCloudflareCustomHostnameFilter struct {
	// Filter by the certificate authority that issued the SSL certificate. Available values: "google", "lets_encrypt", "ssl_com".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#certificate_authority DataCloudflareCustomHostname#certificate_authority}
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Filter by custom origin server name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#custom_origin_server DataCloudflareCustomHostname#custom_origin_server}
	CustomOriginServer *string `field:"optional" json:"customOriginServer" yaml:"customOriginServer"`
	// Direction to order hostnames. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#direction DataCloudflareCustomHostname#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#hostname DataCloudflareCustomHostname#hostname}.
	Hostname *DataCloudflareCustomHostnameFilterHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// Filter by the hostname's activation status.
	//
	// Available values: "active", "pending", "active_redeploying", "moved", "pending_deletion", "deleted", "pending_blocked", "pending_migration", "pending_provisioned", "test_pending", "test_active", "test_active_apex", "test_blocked", "test_failed", "provisioned", "blocked".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#hostname_status DataCloudflareCustomHostname#hostname_status}
	HostnameStatus *string `field:"optional" json:"hostnameStatus" yaml:"hostnameStatus"`
	// Hostname ID to match against.
	//
	// This ID was generated and returned during the initial custom_hostname creation. This parameter cannot be used with the 'hostname' parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#id DataCloudflareCustomHostname#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Field to order hostnames by. Available values: "ssl", "ssl_status".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#order DataCloudflareCustomHostname#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Whether to filter hostnames based on if they have SSL enabled. Available values: 0, 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#ssl DataCloudflareCustomHostname#ssl}
	Ssl *float64 `field:"optional" json:"ssl" yaml:"ssl"`
	// Filter by SSL certificate status.
	//
	// Available values: "initializing", "pending_validation", "deleted", "pending_issuance", "pending_deployment", "pending_deletion", "pending_expiration", "expired", "active", "initializing_timed_out", "validation_timed_out", "issuance_timed_out", "deployment_timed_out", "deletion_timed_out", "pending_cleanup", "staging_deployment", "staging_active", "deactivating", "inactive", "backup_issued", "holding_deployment".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#ssl_status DataCloudflareCustomHostname#ssl_status}
	SslStatus *string `field:"optional" json:"sslStatus" yaml:"sslStatus"`
	// Filter by whether the custom hostname is a wildcard hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostname#wildcard DataCloudflareCustomHostname#wildcard}
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
}

