package googleiamworkloadidentitypool


type GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfig struct {
	// A required mapping of a cloud region to the CA pool resource located in that region used for certificate issuance, adhering to these constraints:  * **Key format:** A supported cloud region name equivalent to the location identifier in the corresponding map entry's value.
	//
	// * **Value format:** A valid CA pool resource path format like:
	// 'projects/{project}/locations/{location}/caPools/{ca_pool}'
	// * **Region Matching:** Workloads are ONLY issued certificates from CA pools within the
	// same region. Also the CA pool region (in value) must match the workload's region (key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_iam_workload_identity_pool#ca_pools GoogleIamWorkloadIdentityPool#ca_pools}
	CaPools *map[string]*string `field:"required" json:"caPools" yaml:"caPools"`
	// Key algorithm to use when generating the key pair.
	//
	// This key pair will be used to create
	// the certificate. If unspecified, this will default to 'ECDSA_P256'.
	//
	// * 'RSA_2048': Specifies RSA with a 2048-bit modulus.
	// * 'RSA_3072': Specifies RSA with a 3072-bit modulus.
	// * 'RSA_4096': Specifies RSA with a 4096-bit modulus.
	// * 'ECDSA_P256': Specifies ECDSA with curve P256.
	// * 'ECDSA_P384': Specifies ECDSA with curve P384. Possible values: ["RSA_2048", "RSA_3072", "RSA_4096", "ECDSA_P256", "ECDSA_P384"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_iam_workload_identity_pool#key_algorithm GoogleIamWorkloadIdentityPool#key_algorithm}
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Lifetime of the workload certificates issued by the CA pool in seconds.
	//
	// Must be between
	// '86400s' (24 hours) to '2592000s' (30 days), ends in the suffix "'s'" (indicating seconds)
	// and is preceded by the number of seconds. If unspecified, this will be defaulted to
	// '86400s' (24 hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_iam_workload_identity_pool#lifetime GoogleIamWorkloadIdentityPool#lifetime}
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
	// Rotation window percentage indicating when certificate rotation should be initiated based on remaining lifetime.
	//
	// Must be between '50' - '80'. If unspecified, this will be defaulted
	// to '50'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_iam_workload_identity_pool#rotation_window_percentage GoogleIamWorkloadIdentityPool#rotation_window_percentage}
	RotationWindowPercentage *float64 `field:"optional" json:"rotationWindowPercentage" yaml:"rotationWindowPercentage"`
}

