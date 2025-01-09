package integrationaction


type IntegrationActionCreateResponders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#id IntegrationAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#type IntegrationAction#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

