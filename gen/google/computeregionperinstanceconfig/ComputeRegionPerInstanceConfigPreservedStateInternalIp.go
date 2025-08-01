package computeregionperinstanceconfig


type ComputeRegionPerInstanceConfigPreservedStateInternalIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_per_instance_config#interface_name ComputeRegionPerInstanceConfig#interface_name}.
	InterfaceName *string `field:"required" json:"interfaceName" yaml:"interfaceName"`
	// These stateful IPs will never be released during autohealing, update or VM instance recreate operations.
	//
	// This flag is used to configure if the IP reservation should be deleted after it is no longer used by the group, e.g. when the given instance or the whole group is deleted. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_per_instance_config#auto_delete ComputeRegionPerInstanceConfig#auto_delete}
	AutoDelete *string `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// ip_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_per_instance_config#ip_address ComputeRegionPerInstanceConfig#ip_address}
	IpAddress *ComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddress `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

