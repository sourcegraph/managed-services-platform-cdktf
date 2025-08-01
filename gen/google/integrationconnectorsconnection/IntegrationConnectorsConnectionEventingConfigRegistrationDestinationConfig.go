package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#destination IntegrationConnectorsConnection#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Key for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#key IntegrationConnectorsConnection#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
}

