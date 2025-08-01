package dnsrecordset


type DnsRecordSetRoutingPolicyWrrHealthCheckedTargets struct {
	// The Internet IP addresses to be health checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dns_record_set#external_endpoints DnsRecordSet#external_endpoints}
	ExternalEndpoints *[]*string `field:"optional" json:"externalEndpoints" yaml:"externalEndpoints"`
	// internal_load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dns_record_set#internal_load_balancers DnsRecordSet#internal_load_balancers}
	InternalLoadBalancers interface{} `field:"optional" json:"internalLoadBalancers" yaml:"internalLoadBalancers"`
}

