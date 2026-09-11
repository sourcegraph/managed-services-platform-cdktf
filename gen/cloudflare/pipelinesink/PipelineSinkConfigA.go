package pipelinesink


type PipelineSinkConfigA struct {
	// Cloudflare Account ID for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#account_id PipelineSink#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// R2 Bucket to write to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#bucket PipelineSink#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#credentials PipelineSink#credentials}.
	Credentials *PipelineSinkConfigCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Controls filename prefix/suffix and strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#file_naming PipelineSink#file_naming}
	FileNaming *PipelineSinkConfigFileNaming `field:"optional" json:"fileNaming" yaml:"fileNaming"`
	// Jurisdiction this bucket is hosted in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#jurisdiction PipelineSink#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
	// Table namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#namespace PipelineSink#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Data-layout partitioning for sinks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#partitioning PipelineSink#partitioning}
	Partitioning *PipelineSinkConfigPartitioning `field:"optional" json:"partitioning" yaml:"partitioning"`
	// Subpath within the bucket to write to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#path PipelineSink#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#rolling_policy PipelineSink#rolling_policy}
	RollingPolicy *PipelineSinkConfigRollingPolicy `field:"optional" json:"rollingPolicy" yaml:"rollingPolicy"`
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#table_name PipelineSink#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Authentication token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#token PipelineSink#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

