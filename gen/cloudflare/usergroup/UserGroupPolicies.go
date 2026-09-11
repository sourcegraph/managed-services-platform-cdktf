package usergroup


type UserGroupPolicies struct {
	// Allow or deny operations against the resources. Available values: "allow", "deny".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/user_group#access UserGroup#access}
	Access *string `field:"required" json:"access" yaml:"access"`
	// A set of permission groups that are specified to the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/user_group#permission_groups UserGroup#permission_groups}
	PermissionGroups interface{} `field:"required" json:"permissionGroups" yaml:"permissionGroups"`
	// A set of resource groups that are specified to the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/user_group#resource_groups UserGroup#resource_groups}
	ResourceGroups interface{} `field:"required" json:"resourceGroups" yaml:"resourceGroups"`
}

