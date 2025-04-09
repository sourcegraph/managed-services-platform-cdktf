package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredential struct {
	// Required. A SecretManager resource containing the OAuth token that authorizes the connection. Format: 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_developer_connect_connection#oauth_token_secret_version GoogleDeveloperConnectConnection#oauth_token_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	OauthTokenSecretVersion *string `field:"required" json:"oauthTokenSecretVersion" yaml:"oauthTokenSecretVersion"`
}

