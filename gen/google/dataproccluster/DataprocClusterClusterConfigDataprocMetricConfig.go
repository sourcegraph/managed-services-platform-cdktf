package dataproccluster


type DataprocClusterClusterConfigDataprocMetricConfig struct {
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_cluster#metrics DataprocCluster#metrics}
	Metrics interface{} `field:"required" json:"metrics" yaml:"metrics"`
}

