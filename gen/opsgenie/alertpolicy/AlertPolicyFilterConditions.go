package alertpolicy


type AlertPolicyFilterConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#field AlertPolicy#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#operation AlertPolicy#operation}.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// User defined value that will be compared with alert field according to the operation. Default value is empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#expected_value AlertPolicy#expected_value}
	ExpectedValue *string `field:"optional" json:"expectedValue" yaml:"expectedValue"`
	// If 'field' is set as 'extra-properties', key could be used for key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#key AlertPolicy#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Indicates behaviour of the given operation. Default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#not AlertPolicy#not}
	Not interface{} `field:"optional" json:"not" yaml:"not"`
	// Order of the condition in conditions list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#order AlertPolicy#order}
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

