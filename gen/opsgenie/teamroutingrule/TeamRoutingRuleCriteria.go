package teamroutingrule


type TeamRoutingRuleCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/team_routing_rule#type TeamRoutingRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/team_routing_rule#conditions TeamRoutingRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

