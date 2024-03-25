package googlenetappvolume


type GoogleNetappVolumeSnapshotPolicyDailySchedule struct {
	// The maximum number of snapshots to keep for the daily schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_netapp_volume#snapshots_to_keep GoogleNetappVolume#snapshots_to_keep}
	SnapshotsToKeep *float64 `field:"required" json:"snapshotsToKeep" yaml:"snapshotsToKeep"`
	// Set the hour to create the snapshot (0-23), defaults to midnight (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_netapp_volume#hour GoogleNetappVolume#hour}
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_netapp_volume#minute GoogleNetappVolume#minute}
	Minute *float64 `field:"optional" json:"minute" yaml:"minute"`
}

