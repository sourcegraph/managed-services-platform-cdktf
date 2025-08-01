package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceScheduledBackup struct {
	// A Cloud Storage URI of a folder, in the format gs://<bucket_name>/<path_inside_bucket>.
	//
	// A sub-folder <backup_folder> containing backup files will be stored below it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#backup_location GoogleDataprocMetastoreService#backup_location}
	BackupLocation *string `field:"required" json:"backupLocation" yaml:"backupLocation"`
	// The scheduled interval in Cron format, see https://en.wikipedia.org/wiki/Cron The default is empty: scheduled backup is not enabled. Must be specified to enable scheduled backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#cron_schedule GoogleDataprocMetastoreService#cron_schedule}
	CronSchedule *string `field:"optional" json:"cronSchedule" yaml:"cronSchedule"`
	// Defines whether the scheduled backup is enabled. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#enabled GoogleDataprocMetastoreService#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the time zone to be used when interpreting cronSchedule.
	//
	// Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones), e.g. America/Los_Angeles or Africa/Abidjan. If left unspecified, the default is UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#time_zone GoogleDataprocMetastoreService#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

