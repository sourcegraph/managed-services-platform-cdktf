package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigSelectedApplications struct {
	// namespaced_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#namespaced_names GkeBackupRestorePlan#namespaced_names}
	NamespacedNames interface{} `field:"required" json:"namespacedNames" yaml:"namespacedNames"`
}

