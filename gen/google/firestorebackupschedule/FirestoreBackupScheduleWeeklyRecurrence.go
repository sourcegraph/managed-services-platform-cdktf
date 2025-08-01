package firestorebackupschedule


type FirestoreBackupScheduleWeeklyRecurrence struct {
	// The day of week to run. Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firestore_backup_schedule#day FirestoreBackupSchedule#day}
	Day *string `field:"optional" json:"day" yaml:"day"`
}

