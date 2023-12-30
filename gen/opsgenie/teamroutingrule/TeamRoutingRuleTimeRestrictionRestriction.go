package teamroutingrule


type TeamRoutingRuleTimeRestrictionRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/team_routing_rule#end_hour TeamRoutingRule#end_hour}.
	EndHour *float64 `field:"required" json:"endHour" yaml:"endHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/team_routing_rule#end_min TeamRoutingRule#end_min}.
	EndMin *float64 `field:"required" json:"endMin" yaml:"endMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/team_routing_rule#start_hour TeamRoutingRule#start_hour}.
	StartHour *float64 `field:"required" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/team_routing_rule#start_min TeamRoutingRule#start_min}.
	StartMin *float64 `field:"required" json:"startMin" yaml:"startMin"`
}

