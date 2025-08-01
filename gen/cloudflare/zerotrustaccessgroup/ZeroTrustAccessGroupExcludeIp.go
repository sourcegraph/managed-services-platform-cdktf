package zerotrustaccessgroup


type ZeroTrustAccessGroupExcludeIp struct {
	// An IPv4 or IPv6 CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#ip ZeroTrustAccessGroup#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

