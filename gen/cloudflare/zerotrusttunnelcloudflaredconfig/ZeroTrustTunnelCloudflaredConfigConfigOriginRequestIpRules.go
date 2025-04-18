package zerotrusttunnelcloudflaredconfig


type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestIpRules struct {
	// Whether to allow the IP prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_tunnel_cloudflared_config#allow ZeroTrustTunnelCloudflaredConfigA#allow}
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Ports to use within the IP rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_tunnel_cloudflared_config#ports ZeroTrustTunnelCloudflaredConfigA#ports}
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
	// IP rule prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_tunnel_cloudflared_config#prefix ZeroTrustTunnelCloudflaredConfigA#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

