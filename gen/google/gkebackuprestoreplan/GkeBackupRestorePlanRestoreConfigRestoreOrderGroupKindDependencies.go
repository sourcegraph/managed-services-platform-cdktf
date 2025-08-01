package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigRestoreOrderGroupKindDependencies struct {
	// requiring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#requiring GkeBackupRestorePlan#requiring}
	Requiring *GkeBackupRestorePlanRestoreConfigRestoreOrderGroupKindDependenciesRequiring `field:"required" json:"requiring" yaml:"requiring"`
	// satisfying block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#satisfying GkeBackupRestorePlan#satisfying}
	Satisfying *GkeBackupRestorePlanRestoreConfigRestoreOrderGroupKindDependenciesSatisfying `field:"required" json:"satisfying" yaml:"satisfying"`
}

