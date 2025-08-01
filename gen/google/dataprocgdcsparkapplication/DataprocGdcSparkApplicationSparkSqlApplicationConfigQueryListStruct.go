package dataprocgdcsparkapplication


type DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct struct {
	// The queries to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_gdc_spark_application#queries DataprocGdcSparkApplication#queries}
	Queries *[]*string `field:"required" json:"queries" yaml:"queries"`
}

