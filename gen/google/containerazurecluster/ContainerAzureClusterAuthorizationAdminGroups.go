package containerazurecluster


type ContainerAzureClusterAuthorizationAdminGroups struct {
	// The name of the group, e.g. `my-group@domain.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_azure_cluster#group ContainerAzureCluster#group}
	Group *string `field:"required" json:"group" yaml:"group"`
}

