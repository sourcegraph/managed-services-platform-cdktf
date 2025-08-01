package googleworkbenchinstance


type GoogleWorkbenchInstanceGceSetupShieldedInstanceConfig struct {
	// Optional.
	//
	// Defines whether the VM instance has integrity monitoring
	// enabled. Enables monitoring and attestation of the boot integrity of the VM
	// instance. The attestation is performed against the integrity policy baseline.
	// This baseline is initially derived from the implicitly trusted boot image
	// when the VM instance is created. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#enable_integrity_monitoring GoogleWorkbenchInstance#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Optional.
	//
	// Defines whether the VM instance has Secure Boot enabled.
	// Secure Boot helps ensure that the system only runs authentic software by verifying
	// the digital signature of all boot components, and halting the boot process
	// if signature verification fails. Disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#enable_secure_boot GoogleWorkbenchInstance#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Optional. Defines whether the VM instance has the vTPM enabled. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#enable_vtpm GoogleWorkbenchInstance#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

