package transcoderjob


type TranscoderJobConfigAdBreaks struct {
	// Start time in seconds for the ad break, relative to the output file timeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/transcoder_job#start_time_offset TranscoderJob#start_time_offset}
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

