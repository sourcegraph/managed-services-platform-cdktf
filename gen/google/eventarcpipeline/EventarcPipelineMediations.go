package eventarcpipeline


type EventarcPipelineMediations struct {
	// transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/eventarc_pipeline#transformation EventarcPipeline#transformation}
	Transformation *EventarcPipelineMediationsTransformation `field:"optional" json:"transformation" yaml:"transformation"`
}

