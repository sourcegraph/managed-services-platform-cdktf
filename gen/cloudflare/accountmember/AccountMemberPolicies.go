package accountmember


type AccountMemberPolicies struct {
	// Allow or deny operations against the resources. Available values: "allow", "deny".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_member#access AccountMember#access}
	Access *string `field:"required" json:"access" yaml:"access"`
	// A set of permission groups that are specified to the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_member#permission_groups AccountMember#permission_groups}
	PermissionGroups interface{} `field:"required" json:"permissionGroups" yaml:"permissionGroups"`
	// A list of resource groups that the policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_member#resource_groups AccountMember#resource_groups}
	ResourceGroups interface{} `field:"required" json:"resourceGroups" yaml:"resourceGroups"`
}

