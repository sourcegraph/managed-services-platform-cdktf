package iamworkloadidentitypoolprovider


type IamWorkloadIdentityPoolProviderSaml struct {
	// SAML Identity provider configuration metadata xml doc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iam_workload_identity_pool_provider#idp_metadata_xml IamWorkloadIdentityPoolProvider#idp_metadata_xml}
	IdpMetadataXml *string `field:"required" json:"idpMetadataXml" yaml:"idpMetadataXml"`
}

