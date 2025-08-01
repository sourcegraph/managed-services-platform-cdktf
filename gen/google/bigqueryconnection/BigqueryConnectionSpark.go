package bigqueryconnection


type BigqueryConnectionSpark struct {
	// metastore_service_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_connection#metastore_service_config BigqueryConnection#metastore_service_config}
	MetastoreServiceConfig *BigqueryConnectionSparkMetastoreServiceConfig `field:"optional" json:"metastoreServiceConfig" yaml:"metastoreServiceConfig"`
	// spark_history_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_connection#spark_history_server_config BigqueryConnection#spark_history_server_config}
	SparkHistoryServerConfig *BigqueryConnectionSparkSparkHistoryServerConfig `field:"optional" json:"sparkHistoryServerConfig" yaml:"sparkHistoryServerConfig"`
}

