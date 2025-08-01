package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfigAuthConfigUserPassword struct {
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#password IntegrationConnectorsConnection#password}
	Password *IntegrationConnectorsConnectionEventingConfigAuthConfigUserPasswordPassword `field:"optional" json:"password" yaml:"password"`
	// Username for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#username IntegrationConnectorsConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

