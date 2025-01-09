package googlenetworkservicestcproute


type GoogleNetworkServicesTcpRouteRulesAction struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_services_tcp_route#destinations GoogleNetworkServicesTcpRoute#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Specifies the idle timeout for the selected route.
	//
	// The idle timeout is defined as the period in which there are no bytes sent or received on either the upstream or downstream connection. If not set, the default idle timeout is 30 seconds. If set to 0s, the timeout will be disabled.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_services_tcp_route#idle_timeout GoogleNetworkServicesTcpRoute#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// If true, Router will use the destination IP and port of the original connection as the destination of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_services_tcp_route#original_destination GoogleNetworkServicesTcpRoute#original_destination}
	OriginalDestination interface{} `field:"optional" json:"originalDestination" yaml:"originalDestination"`
}

