package integrationaction


type IntegrationActionAcknowledgeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#type IntegrationAction#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#conditions IntegrationAction#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

