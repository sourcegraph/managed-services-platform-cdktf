package parametermanagerparameter


type ParameterManagerParameterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_parameter#create ParameterManagerParameter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_parameter#delete ParameterManagerParameter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_parameter#update ParameterManagerParameter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

