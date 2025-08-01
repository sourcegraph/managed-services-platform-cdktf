package iamworkloadidentitypoolprovider


type IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchors struct {
	// PEM certificate of the PKI used for validation. Must only contain one ca certificate(either root or intermediate cert).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iam_workload_identity_pool_provider#pem_certificate IamWorkloadIdentityPoolProvider#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
}

