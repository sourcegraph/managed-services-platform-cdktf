package googlevmwareengineprivatecloud


type GoogleVmwareenginePrivateCloudManagementCluster struct {
	// The user-provided identifier of the new Cluster.
	//
	// The identifier must meet the following requirements:
	//   * Only contains 1-63 alphanumeric characters and hyphens
	//   * Begins with an alphabetical character
	//   * Ends with a non-hyphen character
	//   * Not formatted as a UUID
	//   * Complies with RFC 1034 (https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_vmwareengine_private_cloud#cluster_id GoogleVmwareenginePrivateCloud#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// autoscaling_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_vmwareengine_private_cloud#autoscaling_settings GoogleVmwareenginePrivateCloud#autoscaling_settings}
	AutoscalingSettings *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings `field:"optional" json:"autoscalingSettings" yaml:"autoscalingSettings"`
	// node_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_vmwareengine_private_cloud#node_type_configs GoogleVmwareenginePrivateCloud#node_type_configs}
	NodeTypeConfigs interface{} `field:"optional" json:"nodeTypeConfigs" yaml:"nodeTypeConfigs"`
	// stretched_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_vmwareengine_private_cloud#stretched_cluster_config GoogleVmwareenginePrivateCloud#stretched_cluster_config}
	StretchedClusterConfig *GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfig `field:"optional" json:"stretchedClusterConfig" yaml:"stretchedClusterConfig"`
}

