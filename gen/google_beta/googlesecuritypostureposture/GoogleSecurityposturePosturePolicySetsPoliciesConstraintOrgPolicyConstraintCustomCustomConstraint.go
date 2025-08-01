package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomCustomConstraint struct {
	// The action to take if the condition is met. Possible values: ["ALLOW", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_securityposture_posture#action_type GoogleSecurityposturePosture#action_type}
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// A CEL condition that refers to a supported service resource, for example 'resource.management.autoUpgrade == false'. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_securityposture_posture#condition GoogleSecurityposturePosture#condition}
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// A list of RESTful methods for which to enforce the constraint.
	//
	// Can be 'CREATE', 'UPDATE', or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_securityposture_posture#method_types GoogleSecurityposturePosture#method_types}
	MethodTypes *[]*string `field:"required" json:"methodTypes" yaml:"methodTypes"`
	// Immutable. The name of the custom constraint. This is unique within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_securityposture_posture#name GoogleSecurityposturePosture#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Immutable.
	//
	// The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, 'container.googleapis.com/NodePool'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_securityposture_posture#resource_types GoogleSecurityposturePosture#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// A human-friendly description of the constraint to display as an error message when the policy is violated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_securityposture_posture#description GoogleSecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A human-friendly name for the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_securityposture_posture#display_name GoogleSecurityposturePosture#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

