package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolUpdateSettings struct {
	// surge_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_container_aws_node_pool#surge_settings GoogleContainerAwsNodePool#surge_settings}
	SurgeSettings *GoogleContainerAwsNodePoolUpdateSettingsSurgeSettings `field:"optional" json:"surgeSettings" yaml:"surgeSettings"`
}

