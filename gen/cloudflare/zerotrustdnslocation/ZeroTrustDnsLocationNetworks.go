package zerotrustdnslocation


type ZeroTrustDnsLocationNetworks struct {
	// Specify the IPv4 address or IPv4 CIDR. Limit IPv4 CIDRs to a maximum of /24.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_dns_location#network ZeroTrustDnsLocation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

