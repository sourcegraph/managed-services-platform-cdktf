package integrationaction


type IntegrationActionCloseFilterConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#field IntegrationAction#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#operation IntegrationAction#operation}.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#expected_value IntegrationAction#expected_value}.
	ExpectedValue *string `field:"optional" json:"expectedValue" yaml:"expectedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#key IntegrationAction#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#not IntegrationAction#not}.
	Not interface{} `field:"optional" json:"not" yaml:"not"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#order IntegrationAction#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

