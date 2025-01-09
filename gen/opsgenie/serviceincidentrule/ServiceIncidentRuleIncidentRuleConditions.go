package serviceincidentrule


type ServiceIncidentRuleIncidentRuleConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#field ServiceIncidentRule#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#operation ServiceIncidentRule#operation}.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// User defined value that will be compared with alert field according to the operation. Default value is empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#expected_value ServiceIncidentRule#expected_value}
	ExpectedValue *string `field:"optional" json:"expectedValue" yaml:"expectedValue"`
	// If 'field' is set as 'extra-properties', key could be used for key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#key ServiceIncidentRule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Indicates behaviour of the given operation. Default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#not ServiceIncidentRule#not}
	Not interface{} `field:"optional" json:"not" yaml:"not"`
}

