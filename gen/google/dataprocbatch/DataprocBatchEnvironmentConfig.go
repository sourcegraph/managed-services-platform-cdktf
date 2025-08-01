package dataprocbatch


type DataprocBatchEnvironmentConfig struct {
	// execution_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_batch#execution_config DataprocBatch#execution_config}
	ExecutionConfig *DataprocBatchEnvironmentConfigExecutionConfig `field:"optional" json:"executionConfig" yaml:"executionConfig"`
	// peripherals_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_batch#peripherals_config DataprocBatch#peripherals_config}
	PeripheralsConfig *DataprocBatchEnvironmentConfigPeripheralsConfig `field:"optional" json:"peripheralsConfig" yaml:"peripheralsConfig"`
}

