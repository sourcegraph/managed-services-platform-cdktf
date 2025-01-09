package googletranscoderjob


type GoogleTranscoderJobConfigEncryptionsMpegCenc struct {
	// Specify the encryption scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_transcoder_job#scheme GoogleTranscoderJob#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
}

