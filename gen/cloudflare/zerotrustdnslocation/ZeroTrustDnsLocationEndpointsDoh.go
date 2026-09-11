package zerotrustdnslocation


type ZeroTrustDnsLocationEndpointsDoh struct {
	// Indicate whether the DOH endpoint is enabled for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#enabled ZeroTrustDnsLocation#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the list of allowed source IP network ranges for this endpoint.
	//
	// When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#networks ZeroTrustDnsLocation#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
	// Specify whether the DOH endpoint requires user identity authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#require_token ZeroTrustDnsLocation#require_token}
	RequireToken interface{} `field:"optional" json:"requireToken" yaml:"requireToken"`
}

