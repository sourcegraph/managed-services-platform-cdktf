package googlecomputeinstance


type GoogleComputeInstanceBootDiskInitializeParams struct {
	// The architecture of the disk. One of "X86_64" or "ARM64".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#architecture GoogleComputeInstance#architecture}
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// A flag to enable confidential compute mode on boot disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#enable_confidential_compute GoogleComputeInstance#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
	// The image from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#image GoogleComputeInstance#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// A set of key/value label pairs assigned to the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#labels GoogleComputeInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Indicates how many IOPS to provision for the disk.
	//
	// This sets the number of I/O operations per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#provisioned_iops GoogleComputeInstance#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Indicates how much throughput to provision for the disk.
	//
	// This sets the number of throughput mb per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#provisioned_throughput GoogleComputeInstance#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#resource_manager_tags GoogleComputeInstance#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// A list of self_links of resource policies to attach to the instance's boot disk.
	//
	// Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#resource_policies GoogleComputeInstance#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// The size of the image in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#size GoogleComputeInstance#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The snapshot from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#snapshot GoogleComputeInstance#snapshot}
	Snapshot *string `field:"optional" json:"snapshot" yaml:"snapshot"`
	// source_image_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#source_image_encryption_key GoogleComputeInstance#source_image_encryption_key}
	SourceImageEncryptionKey *GoogleComputeInstanceBootDiskInitializeParamsSourceImageEncryptionKey `field:"optional" json:"sourceImageEncryptionKey" yaml:"sourceImageEncryptionKey"`
	// source_snapshot_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#source_snapshot_encryption_key GoogleComputeInstance#source_snapshot_encryption_key}
	SourceSnapshotEncryptionKey *GoogleComputeInstanceBootDiskInitializeParamsSourceSnapshotEncryptionKey `field:"optional" json:"sourceSnapshotEncryptionKey" yaml:"sourceSnapshotEncryptionKey"`
	// The URL of the storage pool in which the new disk is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#storage_pool GoogleComputeInstance#storage_pool}
	StoragePool *string `field:"optional" json:"storagePool" yaml:"storagePool"`
	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_instance#type GoogleComputeInstance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

