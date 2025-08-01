package transcoderjob


type TranscoderJobConfigElementaryStreamsVideoStream struct {
	// h264 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/transcoder_job#h264 TranscoderJob#h264}
	H264 *TranscoderJobConfigElementaryStreamsVideoStreamH264 `field:"optional" json:"h264" yaml:"h264"`
}

