package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues struct {
	// List of values allowed at this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#allowed_values SecurityposturePosture#allowed_values}
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// List of values denied at this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#denied_values SecurityposturePosture#denied_values}
	DeniedValues *[]*string `field:"optional" json:"deniedValues" yaml:"deniedValues"`
}

