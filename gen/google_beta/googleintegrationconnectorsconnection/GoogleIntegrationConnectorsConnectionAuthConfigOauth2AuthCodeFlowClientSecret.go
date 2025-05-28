package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret struct {
	// The resource name of the secret version in the format, format as: projects/* /secrets/* /versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_integration_connectors_connection#secret_version GoogleIntegrationConnectorsConnection#secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

