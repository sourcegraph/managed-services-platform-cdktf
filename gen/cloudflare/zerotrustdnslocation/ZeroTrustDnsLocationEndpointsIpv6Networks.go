package zerotrustdnslocation


type ZeroTrustDnsLocationEndpointsIpv6Networks struct {
	// Specify the IPv6 address or IPv6 CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_dns_location#network ZeroTrustDnsLocation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

