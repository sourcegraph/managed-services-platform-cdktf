package teamroutingrule


type TeamRoutingRuleTimeRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#type TeamRoutingRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#restriction TeamRoutingRule#restriction}
	Restriction *TeamRoutingRuleTimeRestrictionRestriction `field:"optional" json:"restriction" yaml:"restriction"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#restrictions TeamRoutingRule#restrictions}
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
}

