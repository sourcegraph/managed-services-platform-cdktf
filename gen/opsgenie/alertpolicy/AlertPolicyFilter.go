package alertpolicy


type AlertPolicyFilter struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#conditions AlertPolicy#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#type AlertPolicy#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

