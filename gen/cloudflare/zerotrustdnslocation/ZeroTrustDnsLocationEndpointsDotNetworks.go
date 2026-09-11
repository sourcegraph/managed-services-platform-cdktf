package zerotrustdnslocation


type ZeroTrustDnsLocationEndpointsDotNetworks struct {
	// Specify the IP address or IP CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#network ZeroTrustDnsLocation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

