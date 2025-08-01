package accounttoken


type AccountTokenPolicies struct {
	// Allow or deny operations against the resources. Available values: "allow", "deny".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_token#effect AccountToken#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// A set of permission groups that are specified to the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_token#permission_groups AccountToken#permission_groups}
	PermissionGroups interface{} `field:"required" json:"permissionGroups" yaml:"permissionGroups"`
	// A list of resource names that the policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_token#resources AccountToken#resources}
	Resources *map[string]*string `field:"required" json:"resources" yaml:"resources"`
}

