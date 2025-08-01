package dnspolicy


type DnsPolicyDns64Config struct {
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dns_policy#scope DnsPolicy#scope}
	Scope *DnsPolicyDns64ConfigScope `field:"required" json:"scope" yaml:"scope"`
}

