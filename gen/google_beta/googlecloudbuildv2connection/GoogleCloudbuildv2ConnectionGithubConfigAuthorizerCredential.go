package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGithubConfigAuthorizerCredential struct {
	// A SecretManager resource containing the OAuth token that authorizes the Cloud Build connection. Format: `projects/*\/secrets/*\/versions/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_cloudbuildv2_connection#oauth_token_secret_version GoogleCloudbuildv2Connection#oauth_token_secret_version}
	OauthTokenSecretVersion *string `field:"optional" json:"oauthTokenSecretVersion" yaml:"oauthTokenSecretVersion"`
}

