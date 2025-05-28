package googledataprocgdcsparkapplication


type GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct struct {
	// The queries to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dataproc_gdc_spark_application#queries GoogleDataprocGdcSparkApplication#queries}
	Queries *[]*string `field:"required" json:"queries" yaml:"queries"`
}

