package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigRestoreOrderGroupKindDependenciesRequiring struct {
	// API Group of a Kubernetes resource, e.g. "apiextensions.k8s.io", "storage.k8s.io", etc. Use empty string for core group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#resource_group GkeBackupRestorePlan#resource_group}
	ResourceGroup *string `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// Kind of a Kubernetes resource, e.g. "CustomResourceDefinition", "StorageClass", etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#resource_kind GkeBackupRestorePlan#resource_kind}
	ResourceKind *string `field:"optional" json:"resourceKind" yaml:"resourceKind"`
}

