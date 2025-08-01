package magicwangretunnel


type MagicWanGreTunnelHealthCheck struct {
	// The direction of the flow of the healthcheck.
	//
	// Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel.
	// Available values: "unidirectional", "bidirectional".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_wan_gre_tunnel#direction MagicWanGreTunnel#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_wan_gre_tunnel#enabled MagicWanGreTunnel#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// How frequent the health check is run. The default value is `mid`. Available values: "low", "mid", "high".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_wan_gre_tunnel#rate MagicWanGreTunnel#rate}
	Rate *string `field:"optional" json:"rate" yaml:"rate"`
	// The destination address in a request type health check.
	//
	// After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target. Must be in object form if the x-magic-new-hc-target header is set to true and string form if x-magic-new-hc-target is absent or set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_wan_gre_tunnel#target MagicWanGreTunnel#target}
	Target *MagicWanGreTunnelHealthCheckTarget `field:"optional" json:"target" yaml:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`. Available values: "reply", "request".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_wan_gre_tunnel#type MagicWanGreTunnel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

