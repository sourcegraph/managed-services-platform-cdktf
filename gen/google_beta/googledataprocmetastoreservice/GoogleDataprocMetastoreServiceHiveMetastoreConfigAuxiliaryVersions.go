package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceHiveMetastoreConfigAuxiliaryVersions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#key GoogleDataprocMetastoreService#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Hive metastore version of the auxiliary service. It must be less than the primary Hive metastore service's version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#version GoogleDataprocMetastoreService#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// A mapping of Hive metastore configuration key-value pairs to apply to the auxiliary Hive metastore (configured in hive-site.xml) in addition to the primary version's overrides. If keys are present in both the auxiliary version's overrides and the primary version's overrides, the value from the auxiliary version's overrides takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#config_overrides GoogleDataprocMetastoreService#config_overrides}
	ConfigOverrides *map[string]*string `field:"optional" json:"configOverrides" yaml:"configOverrides"`
}

