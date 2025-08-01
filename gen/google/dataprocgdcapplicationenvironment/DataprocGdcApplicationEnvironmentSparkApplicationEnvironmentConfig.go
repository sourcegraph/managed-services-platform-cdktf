package dataprocgdcapplicationenvironment


type DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig struct {
	// A map of default Spark properties to apply to workloads in this application environment.
	//
	// These defaults may be overridden by per-application properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_gdc_application_environment#default_properties DataprocGdcApplicationEnvironment#default_properties}
	DefaultProperties *map[string]*string `field:"optional" json:"defaultProperties" yaml:"defaultProperties"`
	// The default Dataproc version to use for applications submitted to this application environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_gdc_application_environment#default_version DataprocGdcApplicationEnvironment#default_version}
	DefaultVersion *string `field:"optional" json:"defaultVersion" yaml:"defaultVersion"`
}

