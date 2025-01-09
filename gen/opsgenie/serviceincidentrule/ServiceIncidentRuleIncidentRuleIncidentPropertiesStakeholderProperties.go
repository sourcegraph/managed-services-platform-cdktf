package serviceincidentrule


type ServiceIncidentRuleIncidentRuleIncidentPropertiesStakeholderProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#message ServiceIncidentRule#message}.
	Message *string `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#description ServiceIncidentRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/service_incident_rule#enable ServiceIncidentRule#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

