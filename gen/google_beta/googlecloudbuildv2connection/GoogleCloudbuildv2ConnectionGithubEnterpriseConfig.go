package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGithubEnterpriseConfig struct {
	// Required. The URI of the GitHub Enterprise host this connection is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#host_uri GoogleCloudbuildv2Connection#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// Id of the GitHub App created from the manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#app_id GoogleCloudbuildv2Connection#app_id}
	AppId *float64 `field:"optional" json:"appId" yaml:"appId"`
	// ID of the installation of the GitHub App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#app_installation_id GoogleCloudbuildv2Connection#app_installation_id}
	AppInstallationId *float64 `field:"optional" json:"appInstallationId" yaml:"appInstallationId"`
	// The URL-friendly name of the GitHub App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#app_slug GoogleCloudbuildv2Connection#app_slug}
	AppSlug *string `field:"optional" json:"appSlug" yaml:"appSlug"`
	// SecretManager resource containing the private key of the GitHub App, formatted as `projects/*\/secrets/*\/versions/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#private_key_secret_version GoogleCloudbuildv2Connection#private_key_secret_version}
	PrivateKeySecretVersion *string `field:"optional" json:"privateKeySecretVersion" yaml:"privateKeySecretVersion"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#service_directory_config GoogleCloudbuildv2Connection#service_directory_config}
	ServiceDirectoryConfig *GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// SSL certificate to use for requests to GitHub Enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#ssl_ca GoogleCloudbuildv2Connection#ssl_ca}
	SslCa *string `field:"optional" json:"sslCa" yaml:"sslCa"`
	// SecretManager resource containing the webhook secret of the GitHub App, formatted as `projects/*\/secrets/*\/versions/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudbuildv2_connection#webhook_secret_secret_version GoogleCloudbuildv2Connection#webhook_secret_secret_version}
	WebhookSecretSecretVersion *string `field:"optional" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
}
