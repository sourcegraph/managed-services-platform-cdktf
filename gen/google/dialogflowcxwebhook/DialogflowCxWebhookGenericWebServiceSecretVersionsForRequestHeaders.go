package dialogflowcxwebhook


type DialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_webhook#key DialogflowCxWebhook#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The SecretManager secret version resource storing the header value. Format: 'projects/{project}/secrets/{secret}/versions/{version}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_webhook#secret_version DialogflowCxWebhook#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

