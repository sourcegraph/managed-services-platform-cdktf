package storageanywherecache


type StorageAnywhereCacheTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_anywhere_cache#create StorageAnywhereCache#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_anywhere_cache#delete StorageAnywhereCache#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_anywhere_cache#update StorageAnywhereCache#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

