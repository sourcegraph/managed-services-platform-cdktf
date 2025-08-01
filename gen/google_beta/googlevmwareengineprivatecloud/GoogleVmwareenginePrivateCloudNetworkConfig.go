package googlevmwareengineprivatecloud


type GoogleVmwareenginePrivateCloudNetworkConfig struct {
	// Management CIDR used by VMware management appliances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vmwareengine_private_cloud#management_cidr GoogleVmwareenginePrivateCloud#management_cidr}
	ManagementCidr *string `field:"required" json:"managementCidr" yaml:"managementCidr"`
	// The relative resource name of the VMware Engine network attached to the private cloud.
	//
	// Specify the name in the following form: projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
	// where {project} can either be a project number or a project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vmwareengine_private_cloud#vmware_engine_network GoogleVmwareenginePrivateCloud#vmware_engine_network}
	VmwareEngineNetwork *string `field:"optional" json:"vmwareEngineNetwork" yaml:"vmwareEngineNetwork"`
}

