package zerotrustaccessinfrastructuretarget


type ZeroTrustAccessInfrastructureTargetIpIpv4 struct {
	// IP address of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_infrastructure_target#ip_addr ZeroTrustAccessInfrastructureTarget#ip_addr}
	IpAddr *string `field:"optional" json:"ipAddr" yaml:"ipAddr"`
	// (optional) Private virtual network identifier for the target. If omitted, the default virtual network ID will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_infrastructure_target#virtual_network_id ZeroTrustAccessInfrastructureTarget#virtual_network_id}
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

