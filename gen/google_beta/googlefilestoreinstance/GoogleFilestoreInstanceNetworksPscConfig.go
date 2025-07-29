package googlefilestoreinstance


type GoogleFilestoreInstanceNetworksPscConfig struct {
	// Consumer service project in which the Private Service Connect endpoint would be set up.
	//
	// This is optional, and only relevant in case the network
	// is a shared VPC. If this is not specified, the endpoint would be set up
	// in the VPC host project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_filestore_instance#endpoint_project GoogleFilestoreInstance#endpoint_project}
	EndpointProject *string `field:"optional" json:"endpointProject" yaml:"endpointProject"`
}

