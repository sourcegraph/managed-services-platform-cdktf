package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupScheduleRpoConfig struct {
	// Defines the target RPO for the BackupPlan in minutes, which means the target maximum data loss in time that is acceptable for this BackupPlan.
	//
	// This must be
	// at least 60, i.e., 1 hour, and at most 86400, i.e., 60 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_backup_backup_plan#target_rpo_minutes GoogleGkeBackupBackupPlan#target_rpo_minutes}
	TargetRpoMinutes *float64 `field:"required" json:"targetRpoMinutes" yaml:"targetRpoMinutes"`
	// exclusion_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_backup_backup_plan#exclusion_windows GoogleGkeBackupBackupPlan#exclusion_windows}
	ExclusionWindows interface{} `field:"optional" json:"exclusionWindows" yaml:"exclusionWindows"`
}

