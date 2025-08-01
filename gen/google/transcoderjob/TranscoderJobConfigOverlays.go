package transcoderjob


type TranscoderJobConfigOverlays struct {
	// animations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/transcoder_job#animations TranscoderJob#animations}
	Animations interface{} `field:"optional" json:"animations" yaml:"animations"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/transcoder_job#image TranscoderJob#image}
	Image *TranscoderJobConfigOverlaysImage `field:"optional" json:"image" yaml:"image"`
}

