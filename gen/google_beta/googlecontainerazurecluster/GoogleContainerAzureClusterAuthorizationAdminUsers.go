package googlecontainerazurecluster


type GoogleContainerAzureClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_azure_cluster#username GoogleContainerAzureCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

