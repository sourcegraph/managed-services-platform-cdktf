package incidenttemplate


type IncidentTemplateStakeholderProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/incident_template#message IncidentTemplate#message}.
	Message *string `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/incident_template#description IncidentTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/incident_template#enable IncidentTemplate#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

