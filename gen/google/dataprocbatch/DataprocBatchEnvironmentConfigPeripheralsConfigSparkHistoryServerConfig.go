package dataprocbatch


type DataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig struct {
	// Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_batch#dataproc_cluster DataprocBatch#dataproc_cluster}
	DataprocCluster *string `field:"optional" json:"dataprocCluster" yaml:"dataprocCluster"`
}

