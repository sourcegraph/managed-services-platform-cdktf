package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsDnsResolversIpv4 struct {
	// Specify the IPv4 address of the upstream resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#ip ZeroTrustGatewayPolicy#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Specify a port number to use for the upstream resolver. Defaults to 53 if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#port ZeroTrustGatewayPolicy#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Indicate whether to connect to this resolver over a private network. Must set when vnet_id set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#route_through_private_network ZeroTrustGatewayPolicy#route_through_private_network}
	RouteThroughPrivateNetwork interface{} `field:"optional" json:"routeThroughPrivateNetwork" yaml:"routeThroughPrivateNetwork"`
	// Specify an optional virtual network for this resolver. Uses default virtual network id if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#vnet_id ZeroTrustGatewayPolicy#vnet_id}
	VnetId *string `field:"optional" json:"vnetId" yaml:"vnetId"`
}

