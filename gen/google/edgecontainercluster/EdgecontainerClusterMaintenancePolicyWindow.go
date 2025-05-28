package edgecontainercluster


type EdgecontainerClusterMaintenancePolicyWindow struct {
	// recurring_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/edgecontainer_cluster#recurring_window EdgecontainerCluster#recurring_window}
	RecurringWindow *EdgecontainerClusterMaintenancePolicyWindowRecurringWindow `field:"required" json:"recurringWindow" yaml:"recurringWindow"`
}

