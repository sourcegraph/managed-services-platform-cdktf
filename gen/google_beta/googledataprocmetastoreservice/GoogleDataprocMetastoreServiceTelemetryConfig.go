package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceTelemetryConfig struct {
	// The output format of the Dataproc Metastore service's logs. Default value: "JSON" Possible values: ["LEGACY", "JSON"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_dataproc_metastore_service#log_format GoogleDataprocMetastoreService#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
}

