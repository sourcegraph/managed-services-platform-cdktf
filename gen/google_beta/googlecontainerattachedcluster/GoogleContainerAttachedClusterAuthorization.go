package googlecontainerattachedcluster


type GoogleContainerAttachedClusterAuthorization struct {
	// Users that can perform operations as a cluster admin.
	//
	// A managed
	// ClusterRoleBinding will be created to grant the 'cluster-admin' ClusterRole
	// to the users. Up to ten admin users can be provided.
	//
	// For more info on RBAC, see
	// https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_container_attached_cluster#admin_users GoogleContainerAttachedCluster#admin_users}
	AdminUsers *[]*string `field:"optional" json:"adminUsers" yaml:"adminUsers"`
}

