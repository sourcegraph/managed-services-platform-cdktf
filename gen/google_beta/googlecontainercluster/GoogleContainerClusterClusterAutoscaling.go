package googlecontainercluster


type GoogleContainerClusterClusterAutoscaling struct {
	// auto_provisioning_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#auto_provisioning_defaults GoogleContainerCluster#auto_provisioning_defaults}
	AutoProvisioningDefaults *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaults `field:"optional" json:"autoProvisioningDefaults" yaml:"autoProvisioningDefaults"`
	// The list of Google Compute Engine zones in which the NodePool's nodes can be created by NAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#auto_provisioning_locations GoogleContainerCluster#auto_provisioning_locations}
	AutoProvisioningLocations *[]*string `field:"optional" json:"autoProvisioningLocations" yaml:"autoProvisioningLocations"`
	// Configuration options for the Autoscaling profile feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability when deciding to remove nodes from a cluster.
	//
	// Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults to BALANCED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#autoscaling_profile GoogleContainerCluster#autoscaling_profile}
	AutoscalingProfile *string `field:"optional" json:"autoscalingProfile" yaml:"autoscalingProfile"`
	// Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#resource_limits GoogleContainerCluster#resource_limits}
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

