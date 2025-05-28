package nocodemodule


type NoCodeModuleVariableOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/no_code_module#name NoCodeModule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/no_code_module#options NoCodeModule#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/no_code_module#type NoCodeModule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

