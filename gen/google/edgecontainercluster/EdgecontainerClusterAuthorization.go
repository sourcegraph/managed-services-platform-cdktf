package edgecontainercluster


type EdgecontainerClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#admin_users EdgecontainerCluster#admin_users}
	AdminUsers *EdgecontainerClusterAuthorizationAdminUsers `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

