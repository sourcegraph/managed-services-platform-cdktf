package integrationaction


type IntegrationActionIgnore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#name IntegrationAction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#filter IntegrationAction#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#order IntegrationAction#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#type IntegrationAction#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

