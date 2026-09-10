package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#rdp ZeroTrustAccessApplication#rdp}
	Rdp *ZeroTrustAccessApplicationPoliciesConnectionRulesRdp `field:"optional" json:"rdp" yaml:"rdp"`
	// The SSH-specific rules that define how users may connect to the targets secured by your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#ssh ZeroTrustAccessApplication#ssh}
	Ssh *ZeroTrustAccessApplicationPoliciesConnectionRulesSsh `field:"optional" json:"ssh" yaml:"ssh"`
}

