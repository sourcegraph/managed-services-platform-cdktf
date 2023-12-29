package alertpolicy


type AlertPolicyTimeRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#type AlertPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#restriction AlertPolicy#restriction}
	Restriction *AlertPolicyTimeRestrictionRestriction `field:"optional" json:"restriction" yaml:"restriction"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#restrictions AlertPolicy#restrictions}
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
}

