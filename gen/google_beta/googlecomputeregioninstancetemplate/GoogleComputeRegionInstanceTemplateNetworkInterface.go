package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateNetworkInterface struct {
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#access_config GoogleComputeRegionInstanceTemplate#access_config}
	AccessConfig interface{} `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// alias_ip_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#alias_ip_range GoogleComputeRegionInstanceTemplate#alias_ip_range}
	AliasIpRange interface{} `field:"optional" json:"aliasIpRange" yaml:"aliasIpRange"`
	// The prefix length of the primary internal IPv6 range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#internal_ipv6_prefix_length GoogleComputeRegionInstanceTemplate#internal_ipv6_prefix_length}
	InternalIpv6PrefixLength *float64 `field:"optional" json:"internalIpv6PrefixLength" yaml:"internalIpv6PrefixLength"`
	// ipv6_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#ipv6_access_config GoogleComputeRegionInstanceTemplate#ipv6_access_config}
	Ipv6AccessConfig interface{} `field:"optional" json:"ipv6AccessConfig" yaml:"ipv6AccessConfig"`
	// An IPv6 internal network address for this network interface.
	//
	// If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#ipv6_address GoogleComputeRegionInstanceTemplate#ipv6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The name or self_link of the network to attach this interface to.
	//
	// Use network attribute for Legacy or Auto subnetted networks and subnetwork for custom subnetted networks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#network GoogleComputeRegionInstanceTemplate#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The private IP address to assign to the instance. If empty, the address will be automatically assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#network_ip GoogleComputeRegionInstanceTemplate#network_ip}
	NetworkIp *string `field:"optional" json:"networkIp" yaml:"networkIp"`
	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET, MRDMA, and IRDMA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#nic_type GoogleComputeRegionInstanceTemplate#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// The networking queue count that's specified by users for the network interface.
	//
	// Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#queue_count GoogleComputeRegionInstanceTemplate#queue_count}
	QueueCount *float64 `field:"optional" json:"queueCount" yaml:"queueCount"`
	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not.
	//
	// If not specified, IPV4_ONLY will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#stack_type GoogleComputeRegionInstanceTemplate#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// The name of the subnetwork to attach this interface to.
	//
	// The subnetwork must exist in the same region this instance will be created in. Either network or subnetwork must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#subnetwork GoogleComputeRegionInstanceTemplate#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The ID of the project in which the subnetwork belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_template#subnetwork_project GoogleComputeRegionInstanceTemplate#subnetwork_project}
	SubnetworkProject *string `field:"optional" json:"subnetworkProject" yaml:"subnetworkProject"`
}

