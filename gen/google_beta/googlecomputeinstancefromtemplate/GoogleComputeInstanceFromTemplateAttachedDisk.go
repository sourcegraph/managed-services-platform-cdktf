package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateAttachedDisk struct {
	// The name or self_link of the disk attached to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_instance_from_template#source GoogleComputeInstanceFromTemplate#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Name with which the attached disk is accessible under /dev/disk/by-id/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_instance_from_template#device_name GoogleComputeInstanceFromTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk.
	//
	// Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_instance_from_template#disk_encryption_key_raw GoogleComputeInstanceFromTemplate#disk_encryption_key_raw}
	DiskEncryptionKeyRaw *string `field:"optional" json:"diskEncryptionKeyRaw" yaml:"diskEncryptionKeyRaw"`
	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk.
	//
	// Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_instance_from_template#kms_key_self_link GoogleComputeInstanceFromTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_instance_from_template#mode GoogleComputeInstanceFromTemplate#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

