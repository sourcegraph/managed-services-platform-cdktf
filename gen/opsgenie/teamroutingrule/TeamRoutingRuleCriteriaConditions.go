package teamroutingrule


type TeamRoutingRuleCriteriaConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#field TeamRoutingRule#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#operation TeamRoutingRule#operation}.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#expected_value TeamRoutingRule#expected_value}.
	ExpectedValue *string `field:"optional" json:"expectedValue" yaml:"expectedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#key TeamRoutingRule#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#not TeamRoutingRule#not}.
	Not interface{} `field:"optional" json:"not" yaml:"not"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/team_routing_rule#order TeamRoutingRule#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

