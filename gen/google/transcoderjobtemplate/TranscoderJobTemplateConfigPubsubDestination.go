package transcoderjobtemplate


type TranscoderJobTemplateConfigPubsubDestination struct {
	// The name of the Pub/Sub topic to publish job completion notification to. For example: projects/{project}/topics/{topic}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/transcoder_job_template#topic TranscoderJobTemplate#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

