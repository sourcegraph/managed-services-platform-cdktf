package computesnapshotsettings


type ComputeSnapshotSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_snapshot_settings#create ComputeSnapshotSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_snapshot_settings#delete ComputeSnapshotSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_snapshot_settings#update ComputeSnapshotSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

