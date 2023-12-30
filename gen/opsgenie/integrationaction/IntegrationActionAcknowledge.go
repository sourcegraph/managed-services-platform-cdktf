package integrationaction


type IntegrationActionAcknowledge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#name IntegrationAction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#alias IntegrationAction#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#filter IntegrationAction#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#note IntegrationAction#note}.
	Note *string `field:"optional" json:"note" yaml:"note"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#order IntegrationAction#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#type IntegrationAction#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#user IntegrationAction#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

