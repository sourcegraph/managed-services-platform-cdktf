package googlevertexaidataset


type GoogleVertexAiDatasetEncryptionSpec struct {
	// Required.
	//
	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_dataset#kms_key_name GoogleVertexAiDataset#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

