package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateShieldedInstanceConfig struct {
	// Compare the most recent boot measurements to the integrity policy baseline and return a pair of pass/fail results depending on whether they match or not.
	//
	// Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_region_instance_template#enable_integrity_monitoring GoogleComputeRegionInstanceTemplate#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Verify the digital signature of all boot components, and halt the boot process if signature verification fails.
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_region_instance_template#enable_secure_boot GoogleComputeRegionInstanceTemplate#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Use a virtualized trusted platform module, which is a specialized computer chip you can use to encrypt objects like keys and certificates.
	//
	// Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_region_instance_template#enable_vtpm GoogleComputeRegionInstanceTemplate#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}
