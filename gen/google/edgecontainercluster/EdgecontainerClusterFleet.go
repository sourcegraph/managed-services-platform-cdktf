package edgecontainercluster


type EdgecontainerClusterFleet struct {
	// The name of the Fleet host project where this cluster will be registered. Project names are formatted as 'projects/<project-number>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#project EdgecontainerCluster#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

