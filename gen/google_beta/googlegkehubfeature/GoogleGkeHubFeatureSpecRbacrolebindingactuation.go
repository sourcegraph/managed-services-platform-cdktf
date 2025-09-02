package googlegkehubfeature


type GoogleGkeHubFeatureSpecRbacrolebindingactuation struct {
	// The list of allowed custom roles (ClusterRoles).
	//
	// If a custom role is not part of this list, it cannot be used in a fleet scope RBACRoleBinding. If a custom role in this list is in use, it cannot be removed from the list until the scope RBACRolebindings using it are deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_feature#allowed_custom_roles GoogleGkeHubFeature#allowed_custom_roles}
	AllowedCustomRoles *[]*string `field:"optional" json:"allowedCustomRoles" yaml:"allowedCustomRoles"`
}

