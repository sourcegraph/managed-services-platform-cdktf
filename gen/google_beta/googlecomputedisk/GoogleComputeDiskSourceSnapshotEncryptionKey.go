package googlecomputedisk


type GoogleComputeDiskSourceSnapshotEncryptionKey struct {
	// The self link of the encryption key used to encrypt the disk.
	//
	// Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// ('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have
	// 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_disk#kms_key_self_link GoogleComputeDisk#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// The service account used for the encryption request for the given KMS key.
	//
	// If absent, the Compute Engine Service Agent service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_disk#kms_key_service_account GoogleComputeDisk#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_disk#raw_key GoogleComputeDisk#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
}

