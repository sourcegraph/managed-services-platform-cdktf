package alloydbcluster


type AlloydbClusterAutomatedBackupPolicyTimeBasedRetention struct {
	// The retention period. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/alloydb_cluster#retention_period AlloydbCluster#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

