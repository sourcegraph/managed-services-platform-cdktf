package storagebucket


type StorageBucketHierarchicalNamespace struct {
	// Set this field true to organize bucket with logical file system structure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_bucket#enabled StorageBucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

