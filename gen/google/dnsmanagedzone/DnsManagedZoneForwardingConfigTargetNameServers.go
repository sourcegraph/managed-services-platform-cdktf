package dnsmanagedzone


type DnsManagedZoneForwardingConfigTargetNameServers struct {
	// Fully qualified domain name for the forwarding target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dns_managed_zone#domain_name DnsManagedZone#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Forwarding path for this TargetNameServer.
	//
	// If unset or 'default' Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: ["default", "private"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dns_managed_zone#forwarding_path DnsManagedZone#forwarding_path}
	ForwardingPath *string `field:"optional" json:"forwardingPath" yaml:"forwardingPath"`
	// IPv4 address of a target name server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dns_managed_zone#ipv4_address DnsManagedZone#ipv4_address}
	Ipv4Address *string `field:"optional" json:"ipv4Address" yaml:"ipv4Address"`
}

