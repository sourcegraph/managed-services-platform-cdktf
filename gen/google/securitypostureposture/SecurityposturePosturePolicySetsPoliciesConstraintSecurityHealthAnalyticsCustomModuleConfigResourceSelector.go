package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfigResourceSelector struct {
	// The resource types to run the detector on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/securityposture_posture#resource_types SecurityposturePosture#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

