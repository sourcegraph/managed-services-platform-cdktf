package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigPersistentDirectories struct {
	// gce_pd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_workstations_workstation_config#gce_pd GoogleWorkstationsWorkstationConfigA#gce_pd}
	GcePd *GoogleWorkstationsWorkstationConfigPersistentDirectoriesGcePd `field:"optional" json:"gcePd" yaml:"gcePd"`
	// Location of this directory in the running workstation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_workstations_workstation_config#mount_path GoogleWorkstationsWorkstationConfigA#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

