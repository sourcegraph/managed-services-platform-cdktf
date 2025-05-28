package googlecomputefirewallpolicyrule


type GoogleComputeFirewallPolicyRuleMatch struct {
	// layer4_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#layer4_configs GoogleComputeFirewallPolicyRule#layer4_configs}
	Layer4Configs interface{} `field:"required" json:"layer4Configs" yaml:"layer4Configs"`
	// Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#dest_address_groups GoogleComputeFirewallPolicyRule#dest_address_groups}
	DestAddressGroups *[]*string `field:"optional" json:"destAddressGroups" yaml:"destAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic destination.
	//
	// Maximum number of destination fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#dest_fqdns GoogleComputeFirewallPolicyRule#dest_fqdns}
	DestFqdns *[]*string `field:"optional" json:"destFqdns" yaml:"destFqdns"`
	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#dest_ip_ranges GoogleComputeFirewallPolicyRule#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// Network scope of the traffic destination. Possible values: ["INTERNET", "INTRA_VPC", "NON_INTERNET", "VPC_NETWORKS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#dest_network_scope GoogleComputeFirewallPolicyRule#dest_network_scope}
	DestNetworkScope *string `field:"optional" json:"destNetworkScope" yaml:"destNetworkScope"`
	// Region codes whose IP addresses will be used to match for destination of traffic.
	//
	// Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of dest region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#dest_region_codes GoogleComputeFirewallPolicyRule#dest_region_codes}
	DestRegionCodes *[]*string `field:"optional" json:"destRegionCodes" yaml:"destRegionCodes"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#dest_threat_intelligences GoogleComputeFirewallPolicyRule#dest_threat_intelligences}
	DestThreatIntelligences *[]*string `field:"optional" json:"destThreatIntelligences" yaml:"destThreatIntelligences"`
	// Address groups which should be matched against the traffic source. Maximum number of source address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#src_address_groups GoogleComputeFirewallPolicyRule#src_address_groups}
	SrcAddressGroups *[]*string `field:"optional" json:"srcAddressGroups" yaml:"srcAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic source.
	//
	// Maximum number of source fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#src_fqdns GoogleComputeFirewallPolicyRule#src_fqdns}
	SrcFqdns *[]*string `field:"optional" json:"srcFqdns" yaml:"srcFqdns"`
	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#src_ip_ranges GoogleComputeFirewallPolicyRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
	// Networks of the traffic source. It can be either a full or partial url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#src_networks GoogleComputeFirewallPolicyRule#src_networks}
	SrcNetworks *[]*string `field:"optional" json:"srcNetworks" yaml:"srcNetworks"`
	// Network scope of the traffic source. Possible values: ["INTERNET", "INTRA_VPC", "NON_INTERNET", "VPC_NETWORKS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#src_network_scope GoogleComputeFirewallPolicyRule#src_network_scope}
	SrcNetworkScope *string `field:"optional" json:"srcNetworkScope" yaml:"srcNetworkScope"`
	// Region codes whose IP addresses will be used to match for source of traffic.
	//
	// Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of source region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#src_region_codes GoogleComputeFirewallPolicyRule#src_region_codes}
	SrcRegionCodes *[]*string `field:"optional" json:"srcRegionCodes" yaml:"srcRegionCodes"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_firewall_policy_rule#src_threat_intelligences GoogleComputeFirewallPolicyRule#src_threat_intelligences}
	SrcThreatIntelligences *[]*string `field:"optional" json:"srcThreatIntelligences" yaml:"srcThreatIntelligences"`
}

