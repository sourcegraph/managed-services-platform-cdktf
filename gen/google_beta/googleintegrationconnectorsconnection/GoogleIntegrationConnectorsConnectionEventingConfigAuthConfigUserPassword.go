package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionEventingConfigAuthConfigUserPassword struct {
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_integration_connectors_connection#password GoogleIntegrationConnectorsConnection#password}
	Password *GoogleIntegrationConnectorsConnectionEventingConfigAuthConfigUserPasswordPassword `field:"optional" json:"password" yaml:"password"`
	// Username for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_integration_connectors_connection#username GoogleIntegrationConnectorsConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

