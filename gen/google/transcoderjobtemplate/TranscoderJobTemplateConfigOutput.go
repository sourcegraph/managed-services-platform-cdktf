package transcoderjobtemplate


type TranscoderJobTemplateConfigOutput struct {
	// URI for the output file(s). For example, gs://my-bucket/outputs/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/transcoder_job_template#uri TranscoderJobTemplate#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

