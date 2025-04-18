package zerotrustinfrastructureaccesstarget


type ZeroTrustInfrastructureAccessTargetIpIpv6 struct {
	// The IP address of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_infrastructure_access_target#ip_addr ZeroTrustInfrastructureAccessTarget#ip_addr}
	IpAddr *string `field:"required" json:"ipAddr" yaml:"ipAddr"`
	// The private virtual network identifier for the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_infrastructure_access_target#virtual_network_id ZeroTrustInfrastructureAccessTarget#virtual_network_id}
	VirtualNetworkId *string `field:"required" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

