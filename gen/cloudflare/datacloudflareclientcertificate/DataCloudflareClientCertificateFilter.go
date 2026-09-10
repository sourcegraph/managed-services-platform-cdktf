package datacloudflareclientcertificate


type DataCloudflareClientCertificateFilter struct {
	// Limit to the number of records returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificate#limit DataCloudflareClientCertificate#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Offset the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificate#offset DataCloudflareClientCertificate#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Client Certitifcate Status to filter results by. Available values: "all", "active", "pending_reactivation", "pending_revocation", "revoked".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificate#status DataCloudflareClientCertificate#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

