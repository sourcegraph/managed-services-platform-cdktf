package googleappenginefirewallrule


type GoogleAppEngineFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_firewall_rule#create GoogleAppEngineFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_firewall_rule#delete GoogleAppEngineFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_firewall_rule#update GoogleAppEngineFirewallRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

