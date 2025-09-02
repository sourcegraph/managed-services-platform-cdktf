package googlecomputefirewall


type GoogleComputeFirewallParams struct {
	// Resource manager tags to be bound to the firewall.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456. The field is ignored when empty.
	// The field is immutable and causes resource replacement when mutated. This field is only
	// set at create time and modifying this field after creation will trigger recreation.
	// To apply tags to an existing resource, see the google_tags_tag_binding resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_firewall#resource_manager_tags GoogleComputeFirewall#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

