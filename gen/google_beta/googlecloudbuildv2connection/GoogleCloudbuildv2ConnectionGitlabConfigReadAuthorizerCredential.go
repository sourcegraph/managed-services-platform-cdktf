package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential struct {
	// Required. A SecretManager resource containing the user token that authorizes the Cloud Build connection. Format: 'projects/*\/secrets/*\/versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_cloudbuildv2_connection#user_token_secret_version GoogleCloudbuildv2Connection#user_token_secret_version}
	UserTokenSecretVersion *string `field:"required" json:"userTokenSecretVersion" yaml:"userTokenSecretVersion"`
}

