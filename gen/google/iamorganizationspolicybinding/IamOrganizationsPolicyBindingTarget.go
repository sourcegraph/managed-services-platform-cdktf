package iamorganizationspolicybinding


type IamOrganizationsPolicyBindingTarget struct {
	// Required.
	//
	// Immutable. The resource name of the policy to be bound.
	// The binding parent and policy must belong to the same Organization (or Project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/iam_organizations_policy_binding#principal_set IamOrganizationsPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
}

