package zerotrustdnslocation


type ZeroTrustDnsLocationEndpointsDoh struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_dns_location#enabled ZeroTrustDnsLocation#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_dns_location#networks ZeroTrustDnsLocation#networks}.
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

