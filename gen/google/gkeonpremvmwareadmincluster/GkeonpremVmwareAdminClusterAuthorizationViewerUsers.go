package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterAuthorizationViewerUsers struct {
	// The name of the user, e.g. 'my-gcp-id@gmail.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_admin_cluster#username GkeonpremVmwareAdminCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

