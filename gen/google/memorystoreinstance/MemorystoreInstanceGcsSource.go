package memorystoreinstance


type MemorystoreInstanceGcsSource struct {
	// URIs of the GCS objects to import. Example: gs://bucket1/object1, gs://bucket2/folder2/object2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/memorystore_instance#uris MemorystoreInstance#uris}
	Uris *[]*string `field:"required" json:"uris" yaml:"uris"`
}

