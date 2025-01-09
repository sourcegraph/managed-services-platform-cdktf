package serviceincidentrule


type ServiceIncidentRuleIncidentRuleIncidentProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#message ServiceIncidentRule#message}.
	Message *string `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#priority ServiceIncidentRule#priority}.
	Priority *string `field:"required" json:"priority" yaml:"priority"`
	// stakeholder_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#stakeholder_properties ServiceIncidentRule#stakeholder_properties}
	StakeholderProperties interface{} `field:"required" json:"stakeholderProperties" yaml:"stakeholderProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#description ServiceIncidentRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#details ServiceIncidentRule#details}.
	Details *map[string]*string `field:"optional" json:"details" yaml:"details"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#tags ServiceIncidentRule#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

