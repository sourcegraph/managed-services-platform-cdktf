package googlesccprojectcustommodule


type GoogleSccProjectCustomModuleCustomConfig struct {
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_scc_project_custom_module#predicate GoogleSccProjectCustomModule#predicate}
	Predicate *GoogleSccProjectCustomModuleCustomConfigPredicate `field:"required" json:"predicate" yaml:"predicate"`
	// An explanation of the recommended steps that security teams can take to resolve the detected issue.
	//
	// This explanation is returned with each finding generated by
	// this module in the nextSteps property of the finding JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_scc_project_custom_module#recommendation GoogleSccProjectCustomModule#recommendation}
	Recommendation *string `field:"required" json:"recommendation" yaml:"recommendation"`
	// resource_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_scc_project_custom_module#resource_selector GoogleSccProjectCustomModule#resource_selector}
	ResourceSelector *GoogleSccProjectCustomModuleCustomConfigResourceSelector `field:"required" json:"resourceSelector" yaml:"resourceSelector"`
	// The severity to assign to findings generated by the module. Possible values: ["CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_scc_project_custom_module#severity GoogleSccProjectCustomModule#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// custom_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_scc_project_custom_module#custom_output GoogleSccProjectCustomModule#custom_output}
	CustomOutput *GoogleSccProjectCustomModuleCustomConfigCustomOutput `field:"optional" json:"customOutput" yaml:"customOutput"`
	// Text that describes the vulnerability or misconfiguration that the custom module detects.
	//
	// This explanation is returned with each finding instance to
	// help investigators understand the detected issue. The text must be enclosed in quotation marks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_scc_project_custom_module#description GoogleSccProjectCustomModule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}
