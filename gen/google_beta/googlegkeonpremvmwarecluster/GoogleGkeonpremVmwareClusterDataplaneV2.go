package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterDataplaneV2 struct {
	// Enable advanced networking which requires dataplane_v2_enabled to be set true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gkeonprem_vmware_cluster#advanced_networking GoogleGkeonpremVmwareCluster#advanced_networking}
	AdvancedNetworking interface{} `field:"optional" json:"advancedNetworking" yaml:"advancedNetworking"`
	// Enables Dataplane V2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gkeonprem_vmware_cluster#dataplane_v2_enabled GoogleGkeonpremVmwareCluster#dataplane_v2_enabled}
	DataplaneV2Enabled interface{} `field:"optional" json:"dataplaneV2Enabled" yaml:"dataplaneV2Enabled"`
	// Enable Dataplane V2 for clusters with Windows nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gkeonprem_vmware_cluster#windows_dataplane_v2_enabled GoogleGkeonpremVmwareCluster#windows_dataplane_v2_enabled}
	WindowsDataplaneV2Enabled interface{} `field:"optional" json:"windowsDataplaneV2Enabled" yaml:"windowsDataplaneV2Enabled"`
}

