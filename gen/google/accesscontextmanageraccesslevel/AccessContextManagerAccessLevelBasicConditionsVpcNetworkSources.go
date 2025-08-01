package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelBasicConditionsVpcNetworkSources struct {
	// vpc_subnetwork block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_access_level#vpc_subnetwork AccessContextManagerAccessLevel#vpc_subnetwork}
	VpcSubnetwork *AccessContextManagerAccessLevelBasicConditionsVpcNetworkSourcesVpcSubnetwork `field:"optional" json:"vpcSubnetwork" yaml:"vpcSubnetwork"`
}

