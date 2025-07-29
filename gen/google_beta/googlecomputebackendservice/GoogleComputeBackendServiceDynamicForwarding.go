package googlecomputebackendservice


type GoogleComputeBackendServiceDynamicForwarding struct {
	// ip_port_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#ip_port_selection GoogleComputeBackendService#ip_port_selection}
	IpPortSelection *GoogleComputeBackendServiceDynamicForwardingIpPortSelection `field:"optional" json:"ipPortSelection" yaml:"ipPortSelection"`
}

