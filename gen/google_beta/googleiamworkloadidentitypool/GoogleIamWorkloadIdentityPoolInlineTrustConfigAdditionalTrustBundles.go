package googleiamworkloadidentitypool


type GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundles struct {
	// trust_anchors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_iam_workload_identity_pool#trust_anchors GoogleIamWorkloadIdentityPool#trust_anchors}
	TrustAnchors interface{} `field:"required" json:"trustAnchors" yaml:"trustAnchors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_iam_workload_identity_pool#trust_domain GoogleIamWorkloadIdentityPool#trust_domain}.
	TrustDomain *string `field:"required" json:"trustDomain" yaml:"trustDomain"`
}

