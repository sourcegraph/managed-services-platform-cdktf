package googleedgecontainercluster


type GoogleEdgecontainerClusterMaintenancePolicyWindow struct {
	// recurring_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_edgecontainer_cluster#recurring_window GoogleEdgecontainerCluster#recurring_window}
	RecurringWindow *GoogleEdgecontainerClusterMaintenancePolicyWindowRecurringWindow `field:"required" json:"recurringWindow" yaml:"recurringWindow"`
}

