package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesExcludeIp struct {
	// An IPv4 or IPv6 CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_application#ip ZeroTrustAccessApplication#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

