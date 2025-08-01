package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsMaintenanceWindow struct {
	// Day of week (1-7), starting on Monday.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#day GoogleSqlDatabaseInstance#day}
	Day *float64 `field:"optional" json:"day" yaml:"day"`
	// Hour of day (0-23), ignored if day not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#hour GoogleSqlDatabaseInstance#hour}
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
	// Receive updates after one week (canary) or after two weeks (stable) or after five weeks (week5) of notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#update_track GoogleSqlDatabaseInstance#update_track}
	UpdateTrack *string `field:"optional" json:"updateTrack" yaml:"updateTrack"`
}

