package computeinstancesettings


type ComputeInstanceSettingsMetadata struct {
	// A metadata key/value items map. The total size of all keys and values must be less than 512KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_settings#items ComputeInstanceSettings#items}
	Items *map[string]*string `field:"optional" json:"items" yaml:"items"`
}

