package apihubcuration


type ApihubCurationEndpoint struct {
	// application_integration_endpoint_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apihub_curation#application_integration_endpoint_details ApihubCuration#application_integration_endpoint_details}
	ApplicationIntegrationEndpointDetails *ApihubCurationEndpointApplicationIntegrationEndpointDetails `field:"required" json:"applicationIntegrationEndpointDetails" yaml:"applicationIntegrationEndpointDetails"`
}

