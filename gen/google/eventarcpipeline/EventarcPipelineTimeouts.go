package eventarcpipeline


type EventarcPipelineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/eventarc_pipeline#create EventarcPipeline#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/eventarc_pipeline#delete EventarcPipeline#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/eventarc_pipeline#update EventarcPipeline#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

