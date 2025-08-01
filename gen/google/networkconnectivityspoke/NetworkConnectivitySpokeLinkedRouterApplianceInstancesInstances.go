package networkconnectivityspoke


type NetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances struct {
	// The IP address on the VM to use for peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_connectivity_spoke#ip_address NetworkConnectivitySpoke#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The URI of the virtual machine resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_connectivity_spoke#virtual_machine NetworkConnectivitySpoke#virtual_machine}
	VirtualMachine *string `field:"required" json:"virtualMachine" yaml:"virtualMachine"`
}

