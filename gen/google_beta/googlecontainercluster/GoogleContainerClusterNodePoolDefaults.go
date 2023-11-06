package googlecontainercluster


type GoogleContainerClusterNodePoolDefaults struct {
	// node_config_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_container_cluster#node_config_defaults GoogleContainerCluster#node_config_defaults}
	NodeConfigDefaults *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults `field:"optional" json:"nodeConfigDefaults" yaml:"nodeConfigDefaults"`
}

