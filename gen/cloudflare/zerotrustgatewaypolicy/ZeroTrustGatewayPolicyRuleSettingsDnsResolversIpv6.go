package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv6 struct {
	// IPv6 address of upstream resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#ip ZeroTrustGatewayPolicy#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#port ZeroTrustGatewayPolicy#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Whether to connect to this resolver over a private network. Must be set when vnet_id is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#route_through_private_network ZeroTrustGatewayPolicy#route_through_private_network}
	RouteThroughPrivateNetwork interface{} `field:"optional" json:"routeThroughPrivateNetwork" yaml:"routeThroughPrivateNetwork"`
	// Optionally specify a virtual network for this resolver. Uses default virtual network id if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#vnet_id ZeroTrustGatewayPolicy#vnet_id}
	VnetId *string `field:"optional" json:"vnetId" yaml:"vnetId"`
}

