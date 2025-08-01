package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceRestoreBackupContext struct {
	// The ID of the backup run to restore from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#backup_run_id GoogleSqlDatabaseInstance#backup_run_id}
	BackupRunId *float64 `field:"required" json:"backupRunId" yaml:"backupRunId"`
	// The ID of the instance that the backup was taken from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#instance_id GoogleSqlDatabaseInstance#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The full project ID of the source instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#project GoogleSqlDatabaseInstance#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

