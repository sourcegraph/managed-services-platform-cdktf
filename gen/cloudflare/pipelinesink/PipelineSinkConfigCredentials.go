package pipelinesink


type PipelineSinkConfigCredentials struct {
	// Cloudflare Account ID for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#access_key_id PipelineSink#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Cloudflare Account ID for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#secret_access_key PipelineSink#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

