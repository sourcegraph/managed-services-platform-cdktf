package googlebeyondcorpapplication


type GoogleBeyondcorpApplicationUpstreamsEgressPolicy struct {
	// Required. List of regions where the application sends traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_beyondcorp_application#regions GoogleBeyondcorpApplication#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
}

