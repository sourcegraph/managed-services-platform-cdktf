package alloydbcluster


type AlloydbClusterRestoreBackupSource struct {
	// The name of the backup that this cluster is restored from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/alloydb_cluster#backup_name AlloydbCluster#backup_name}
	BackupName *string `field:"required" json:"backupName" yaml:"backupName"`
}

