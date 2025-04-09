package googlealloydbcluster


type GoogleAlloydbClusterContinuousBackupConfig struct {
	// Whether continuous backup recovery is enabled. If not set, defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_alloydb_cluster#enabled GoogleAlloydbCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_alloydb_cluster#encryption_config GoogleAlloydbCluster#encryption_config}
	EncryptionConfig *GoogleAlloydbClusterContinuousBackupConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The numbers of days that are eligible to restore from using PITR.
	//
	// To support the entire recovery window, backups and logs are retained for one day more than the recovery window.
	//
	// If not set, defaults to 14 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_alloydb_cluster#recovery_window_days GoogleAlloydbCluster#recovery_window_days}
	RecoveryWindowDays *float64 `field:"optional" json:"recoveryWindowDays" yaml:"recoveryWindowDays"`
}

