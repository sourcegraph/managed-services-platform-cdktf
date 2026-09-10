package pipelinestream


type PipelineStreamHttp struct {
	// Indicates that authentication is required for the HTTP endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#authentication PipelineStream#authentication}
	Authentication interface{} `field:"required" json:"authentication" yaml:"authentication"`
	// Indicates that the HTTP endpoint is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#enabled PipelineStream#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the CORS options for the HTTP endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#cors PipelineStream#cors}
	Cors *PipelineStreamHttpCors `field:"optional" json:"cors" yaml:"cors"`
}

