package computeregioninstancetemplate


type ComputeRegionInstanceTemplateDiskDiskEncryptionKey struct {
	// The self link of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_instance_template#kms_key_self_link ComputeRegionInstanceTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// The service account being used for the encryption request for the given KMS key.
	//
	// If absent, the Compute Engine default service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_instance_template#kms_key_service_account ComputeRegionInstanceTemplate#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
}

