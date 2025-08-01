package containerawscluster


type ContainerAwsClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_aws_cluster#admin_users ContainerAwsCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
	// admin_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_aws_cluster#admin_groups ContainerAwsCluster#admin_groups}
	AdminGroups interface{} `field:"optional" json:"adminGroups" yaml:"adminGroups"`
}

