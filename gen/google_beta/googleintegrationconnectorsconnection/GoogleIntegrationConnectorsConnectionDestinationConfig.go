package googleintegrationconnectorsconnection

type GoogleIntegrationConnectorsConnectionDestinationConfig struct {
	// The key is the destination identifier that is supported by the Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_connection#key GoogleIntegrationConnectorsConnection#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_connection#destination GoogleIntegrationConnectorsConnection#destination}
	Destination any `field:"optional" json:"destination" yaml:"destination"`
}
