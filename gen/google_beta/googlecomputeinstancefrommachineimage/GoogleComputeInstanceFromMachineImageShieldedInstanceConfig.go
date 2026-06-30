package googlecomputeinstancefrommachineimage

type GoogleComputeInstanceFromMachineImageShieldedInstanceConfig struct {
	// Whether integrity monitoring is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_from_machine_image#enable_integrity_monitoring GoogleComputeInstanceFromMachineImage#enable_integrity_monitoring}
	EnableIntegrityMonitoring any `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Whether secure boot is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_from_machine_image#enable_secure_boot GoogleComputeInstanceFromMachineImage#enable_secure_boot}
	EnableSecureBoot any `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Whether the instance uses vTPM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_from_machine_image#enable_vtpm GoogleComputeInstanceFromMachineImage#enable_vtpm}
	EnableVtpm any `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}
