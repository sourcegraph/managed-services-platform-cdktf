package alertpolicy


type AlertPolicyTimeRestrictionRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#end_hour AlertPolicy#end_hour}.
	EndHour *float64 `field:"required" json:"endHour" yaml:"endHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#end_min AlertPolicy#end_min}.
	EndMin *float64 `field:"required" json:"endMin" yaml:"endMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#start_hour AlertPolicy#start_hour}.
	StartHour *float64 `field:"required" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#start_min AlertPolicy#start_min}.
	StartMin *float64 `field:"required" json:"startMin" yaml:"startMin"`
}

