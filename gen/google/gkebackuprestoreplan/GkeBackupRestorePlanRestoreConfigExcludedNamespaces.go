package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigExcludedNamespaces struct {
	// A list of Kubernetes Namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#namespaces GkeBackupRestorePlan#namespaces}
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

