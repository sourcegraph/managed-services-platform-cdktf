package googlerediscluster


type GoogleRedisClusterAutomatedBackupConfigFixedFrequencyScheduleStartTime struct {
	// Hours of a day in 24 hour format.
	//
	// Must be greater than or equal to 0 and typically must be less than or equal to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_redis_cluster#hours GoogleRedisCluster#hours}
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
}

