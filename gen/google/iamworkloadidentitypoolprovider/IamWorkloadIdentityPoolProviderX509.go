package iamworkloadidentitypoolprovider


type IamWorkloadIdentityPoolProviderX509 struct {
	// trust_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iam_workload_identity_pool_provider#trust_store IamWorkloadIdentityPoolProvider#trust_store}
	TrustStore *IamWorkloadIdentityPoolProviderX509TrustStore `field:"required" json:"trustStore" yaml:"trustStore"`
}

