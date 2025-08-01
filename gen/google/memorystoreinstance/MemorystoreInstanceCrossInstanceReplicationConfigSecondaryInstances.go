package memorystoreinstance


type MemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstances struct {
	// The full resource path of the Nth instance in the format: projects/{project}/locations/{region}/instance/{instance-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/memorystore_instance#instance MemorystoreInstance#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
}

