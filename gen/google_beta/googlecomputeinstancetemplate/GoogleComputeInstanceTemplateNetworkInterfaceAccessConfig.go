package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateNetworkInterfaceAccessConfig struct {
	// The IP address that will be 1:1 mapped to the instance's network ip.
	//
	// If not given, one will be generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_instance_template#nat_ip GoogleComputeInstanceTemplate#nat_ip}
	NatIp *string `field:"optional" json:"natIp" yaml:"natIp"`
	// The networking tier used for configuring this instance template.
	//
	// This field can take the following values: PREMIUM, STANDARD, FIXED_STANDARD. If this field is not specified, it is assumed to be PREMIUM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_instance_template#network_tier GoogleComputeInstanceTemplate#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
}

