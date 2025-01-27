package loadbalancerpool


type LoadBalancerPoolOrigins struct {
	// The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/load_balancer_pool#address LoadBalancerPool#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// A human-identifiable name for the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/load_balancer_pool#name LoadBalancerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether this origin is enabled.
	//
	// Disabled origins will not receive traffic and are excluded from health checks. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/load_balancer_pool#enabled LoadBalancerPool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/load_balancer_pool#header LoadBalancerPool#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also belong to the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/load_balancer_pool#virtual_network_id LoadBalancerPool#virtual_network_id}
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// The weight (0.01 - 1.00) of this origin, relative to other origins in the pool. Equal values mean equal weighting. A weight of 0 means traffic will not be sent to this origin, but health is still checked. When [`origin_steering.policy="least_outstanding_requests"`](#policy), weight is used to scale the origin's outstanding requests. When [`origin_steering.policy="least_connections"`](#policy), weight is used to scale the origin's open connections. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/load_balancer_pool#weight LoadBalancerPool#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

