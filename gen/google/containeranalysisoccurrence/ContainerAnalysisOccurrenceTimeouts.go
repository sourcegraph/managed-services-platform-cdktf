package containeranalysisoccurrence


type ContainerAnalysisOccurrenceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_analysis_occurrence#create ContainerAnalysisOccurrence#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_analysis_occurrence#delete ContainerAnalysisOccurrence#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_analysis_occurrence#update ContainerAnalysisOccurrence#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

