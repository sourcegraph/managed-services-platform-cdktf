package bigtabletable


type BigtableTableAutomatedBackupPolicy struct {
	// How frequently automated backups should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigtable_table#frequency BigtableTable#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// How long the automated backups should be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigtable_table#retention_period BigtableTable#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

