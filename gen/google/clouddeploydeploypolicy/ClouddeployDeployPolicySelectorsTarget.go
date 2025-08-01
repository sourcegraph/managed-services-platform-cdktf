package clouddeploydeploypolicy


type ClouddeployDeployPolicySelectorsTarget struct {
	// ID of the 'Target'.
	//
	// The value of this field could be one of the following: * The last segment of a target name. It only needs the ID to determine which target is being referred to * "*", all targets in a location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_deploy_policy#id ClouddeployDeployPolicy#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_deploy_policy#labels ClouddeployDeployPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

