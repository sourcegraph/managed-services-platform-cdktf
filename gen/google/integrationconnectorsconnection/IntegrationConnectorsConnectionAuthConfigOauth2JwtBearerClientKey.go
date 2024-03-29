package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfigOauth2JwtBearerClientKey struct {
	// The resource name of the secret version in the format, format as: projects/*\/secrets/*\/versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/integration_connectors_connection#secret_version IntegrationConnectorsConnection#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

