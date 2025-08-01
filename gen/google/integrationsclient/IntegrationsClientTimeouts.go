package integrationsclient


type IntegrationsClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integrations_client#create IntegrationsClient#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integrations_client#delete IntegrationsClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

