package pipelinesink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineSinkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Defines the name of the Sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#name PipelineSink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the type of sink. Available values: "r2", "r2_data_catalog".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#type PipelineSink#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the public ID of the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#account_id PipelineSink#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Defines the configuration of the R2 Sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#config PipelineSink#config}
	Config *PipelineSinkConfigA `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#format PipelineSink#format}.
	Format *PipelineSinkFormat `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#schema PipelineSink#schema}.
	Schema *PipelineSinkSchema `field:"optional" json:"schema" yaml:"schema"`
}

