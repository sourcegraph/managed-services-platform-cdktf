package googleosconfigospolicyassignment

type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroups struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_os_policy_assignment#resources GoogleOsConfigOsPolicyAssignment#resources}
	Resources any `field:"required" json:"resources" yaml:"resources"`
	// inventory_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_os_policy_assignment#inventory_filters GoogleOsConfigOsPolicyAssignment#inventory_filters}
	InventoryFilters any `field:"optional" json:"inventoryFilters" yaml:"inventoryFilters"`
}
