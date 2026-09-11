package zerotrustaccessaicontrolsmcpportal


type ZeroTrustAccessAiControlsMcpPortalServersUpdatedPrompts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#name ZeroTrustAccessAiControlsMcpPortal#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#alias ZeroTrustAccessAiControlsMcpPortal#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#description ZeroTrustAccessAiControlsMcpPortal#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_portal#enabled ZeroTrustAccessAiControlsMcpPortal#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

