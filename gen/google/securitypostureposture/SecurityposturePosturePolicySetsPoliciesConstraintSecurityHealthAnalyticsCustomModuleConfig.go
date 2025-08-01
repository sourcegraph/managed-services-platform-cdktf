package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfig struct {
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#predicate SecurityposturePosture#predicate}
	Predicate *SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfigPredicate `field:"required" json:"predicate" yaml:"predicate"`
	// resource_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#resource_selector SecurityposturePosture#resource_selector}
	ResourceSelector *SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfigResourceSelector `field:"required" json:"resourceSelector" yaml:"resourceSelector"`
	// The severity to assign to findings generated by the module. Possible values: ["SEVERITY_UNSPECIFIED", "CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#severity SecurityposturePosture#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// custom_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#custom_output SecurityposturePosture#custom_output}
	CustomOutput *SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutput `field:"optional" json:"customOutput" yaml:"customOutput"`
	// Text that describes the vulnerability or misconfiguration that the custom module detects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#description SecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An explanation of the recommended steps that security teams can take to resolve the detected issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#recommendation SecurityposturePosture#recommendation}
	Recommendation *string `field:"optional" json:"recommendation" yaml:"recommendation"`
}

