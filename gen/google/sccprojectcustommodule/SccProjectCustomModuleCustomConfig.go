package sccprojectcustommodule


type SccProjectCustomModuleCustomConfig struct {
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/scc_project_custom_module#predicate SccProjectCustomModule#predicate}
	Predicate *SccProjectCustomModuleCustomConfigPredicate `field:"required" json:"predicate" yaml:"predicate"`
	// An explanation of the recommended steps that security teams can take to resolve the detected issue.
	//
	// This explanation is returned with each finding generated by
	// this module in the nextSteps property of the finding JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/scc_project_custom_module#recommendation SccProjectCustomModule#recommendation}
	Recommendation *string `field:"required" json:"recommendation" yaml:"recommendation"`
	// resource_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/scc_project_custom_module#resource_selector SccProjectCustomModule#resource_selector}
	ResourceSelector *SccProjectCustomModuleCustomConfigResourceSelector `field:"required" json:"resourceSelector" yaml:"resourceSelector"`
	// The severity to assign to findings generated by the module. Possible values: ["CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/scc_project_custom_module#severity SccProjectCustomModule#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// custom_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/scc_project_custom_module#custom_output SccProjectCustomModule#custom_output}
	CustomOutput *SccProjectCustomModuleCustomConfigCustomOutput `field:"optional" json:"customOutput" yaml:"customOutput"`
	// Text that describes the vulnerability or misconfiguration that the custom module detects.
	//
	// This explanation is returned with each finding instance to
	// help investigators understand the detected issue. The text must be enclosed in quotation marks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/scc_project_custom_module#description SccProjectCustomModule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}
