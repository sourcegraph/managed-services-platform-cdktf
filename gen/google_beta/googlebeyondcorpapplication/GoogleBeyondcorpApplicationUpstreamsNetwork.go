package googlebeyondcorpapplication


type GoogleBeyondcorpApplicationUpstreamsNetwork struct {
	// Required. Network name is of the format: 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_beyondcorp_application#name GoogleBeyondcorpApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

