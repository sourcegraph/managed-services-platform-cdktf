package googlenetworkservicestcproute


type GoogleNetworkServicesTcpRouteRulesMatches struct {
	// Must be specified in the CIDR range format.
	//
	// A CIDR range consists of an IP Address and a prefix length to construct the subnet mask.
	// By default, the prefix length is 32 (i.e. matches a single IP address). Only IPV4 addresses are supported. Examples: "10.0.0.1" - matches against this exact IP address. "10.0.0.0/8" - matches against any IP address within the 10.0.0.0 subnet and 255.255.255.0 mask. "0.0.0.0/0" - matches against any IP address'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_tcp_route#address GoogleNetworkServicesTcpRoute#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Specifies the destination port to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_tcp_route#port GoogleNetworkServicesTcpRoute#port}
	Port *string `field:"required" json:"port" yaml:"port"`
}

