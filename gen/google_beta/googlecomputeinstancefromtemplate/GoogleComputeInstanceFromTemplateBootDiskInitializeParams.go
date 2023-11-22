package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateBootDiskInitializeParams struct {
	// A flag to enable confidential compute mode on boot disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_instance_from_template#enable_confidential_compute GoogleComputeInstanceFromTemplate#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
	// The image from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_instance_from_template#image GoogleComputeInstanceFromTemplate#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// A set of key/value label pairs assigned to the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_instance_from_template#labels GoogleComputeInstanceFromTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_instance_from_template#resource_manager_tags GoogleComputeInstanceFromTemplate#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// The size of the image in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_instance_from_template#size GoogleComputeInstanceFromTemplate#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_instance_from_template#type GoogleComputeInstanceFromTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

