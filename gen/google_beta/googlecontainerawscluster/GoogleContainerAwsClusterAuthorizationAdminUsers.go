package googlecontainerawscluster


type GoogleContainerAwsClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_container_aws_cluster#username GoogleContainerAwsCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

