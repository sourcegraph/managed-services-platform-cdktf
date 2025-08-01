package googlecontainercluster


type GoogleContainerClusterNodePool struct {
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#autoscaling GoogleContainerCluster#autoscaling}
	Autoscaling *GoogleContainerClusterNodePoolAutoscaling `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// The initial number of nodes for the pool.
	//
	// In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#initial_node_count GoogleContainerCluster#initial_node_count}
	InitialNodeCount *float64 `field:"optional" json:"initialNodeCount" yaml:"initialNodeCount"`
	// management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#management GoogleContainerCluster#management}
	Management *GoogleContainerClusterNodePoolManagement `field:"optional" json:"management" yaml:"management"`
	// The maximum number of pods per node in this node pool.
	//
	// Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#max_pods_per_node GoogleContainerCluster#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
	// The name of the node pool. If left blank, Terraform will auto-generate a unique name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#name GoogleContainerCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#name_prefix GoogleContainerCluster#name_prefix}
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#network_config GoogleContainerCluster#network_config}
	NetworkConfig *GoogleContainerClusterNodePoolNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#node_config GoogleContainerCluster#node_config}
	NodeConfig *GoogleContainerClusterNodePoolNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// The number of nodes per instance group.
	//
	// This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#node_count GoogleContainerCluster#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// The list of zones in which the node pool's nodes should be located.
	//
	// Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#node_locations GoogleContainerCluster#node_locations}
	NodeLocations *[]*string `field:"optional" json:"nodeLocations" yaml:"nodeLocations"`
	// placement_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#placement_policy GoogleContainerCluster#placement_policy}
	PlacementPolicy *GoogleContainerClusterNodePoolPlacementPolicy `field:"optional" json:"placementPolicy" yaml:"placementPolicy"`
	// queued_provisioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#queued_provisioning GoogleContainerCluster#queued_provisioning}
	QueuedProvisioning *GoogleContainerClusterNodePoolQueuedProvisioning `field:"optional" json:"queuedProvisioning" yaml:"queuedProvisioning"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#upgrade_settings GoogleContainerCluster#upgrade_settings}
	UpgradeSettings *GoogleContainerClusterNodePoolUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
	// The Kubernetes version for the nodes in this pool.
	//
	// Note that if this field and auto_upgrade are both specified, they will fight each other for what the node version should be, so setting both is highly discouraged. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#version GoogleContainerCluster#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

