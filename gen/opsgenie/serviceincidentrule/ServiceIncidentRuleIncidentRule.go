package serviceincidentrule


type ServiceIncidentRuleIncidentRule struct {
	// incident_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#incident_properties ServiceIncidentRule#incident_properties}
	IncidentProperties interface{} `field:"required" json:"incidentProperties" yaml:"incidentProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#condition_match_type ServiceIncidentRule#condition_match_type}.
	ConditionMatchType *string `field:"optional" json:"conditionMatchType" yaml:"conditionMatchType"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#conditions ServiceIncidentRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

