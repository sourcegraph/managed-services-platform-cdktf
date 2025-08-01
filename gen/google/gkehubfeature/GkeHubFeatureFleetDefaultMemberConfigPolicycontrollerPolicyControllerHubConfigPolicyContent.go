package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent struct {
	// bundles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#bundles GkeHubFeature#bundles}
	Bundles interface{} `field:"optional" json:"bundles" yaml:"bundles"`
	// template_library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#template_library GkeHubFeature#template_library}
	TemplateLibrary *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary `field:"optional" json:"templateLibrary" yaml:"templateLibrary"`
}

