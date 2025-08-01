package integrationconnectorsconnection


type IntegrationConnectorsConnectionLockConfig struct {
	// Indicates whether or not the connection is locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#locked IntegrationConnectorsConnection#locked}
	Locked interface{} `field:"required" json:"locked" yaml:"locked"`
	// Describes why a connection is locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#reason IntegrationConnectorsConnection#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

