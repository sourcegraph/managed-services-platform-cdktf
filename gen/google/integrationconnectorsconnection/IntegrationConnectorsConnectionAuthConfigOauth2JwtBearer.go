package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfigOauth2JwtBearer struct {
	// client_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#client_key IntegrationConnectorsConnection#client_key}
	ClientKey *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearerClientKey `field:"optional" json:"clientKey" yaml:"clientKey"`
	// jwt_claims block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/integration_connectors_connection#jwt_claims IntegrationConnectorsConnection#jwt_claims}
	JwtClaims *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearerJwtClaims `field:"optional" json:"jwtClaims" yaml:"jwtClaims"`
}

