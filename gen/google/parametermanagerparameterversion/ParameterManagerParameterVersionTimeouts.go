package parametermanagerparameterversion


type ParameterManagerParameterVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_parameter_version#create ParameterManagerParameterVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_parameter_version#delete ParameterManagerParameterVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/parameter_manager_parameter_version#update ParameterManagerParameterVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

