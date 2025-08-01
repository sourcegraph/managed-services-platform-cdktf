package googlecloudschedulerjob


type GoogleCloudSchedulerJobHttpTarget struct {
	// The full URI path that the request will be sent to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#uri GoogleCloudSchedulerJob#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// HTTP request body.
	//
	// A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
	// It is an error to set body on a job with an incompatible HttpMethod.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#body GoogleCloudSchedulerJob#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// This map contains the header field names and values.
	//
	// Repeated headers are not supported, but a header value can contain commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#headers GoogleCloudSchedulerJob#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Which HTTP method to use for the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#http_method GoogleCloudSchedulerJob#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// oauth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#oauth_token GoogleCloudSchedulerJob#oauth_token}
	OauthToken *GoogleCloudSchedulerJobHttpTargetOauthToken `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// oidc_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#oidc_token GoogleCloudSchedulerJob#oidc_token}
	OidcToken *GoogleCloudSchedulerJobHttpTargetOidcToken `field:"optional" json:"oidcToken" yaml:"oidcToken"`
}

