package eventarcpipeline


type EventarcPipelineDestinationsAuthenticationConfig struct {
	// google_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/eventarc_pipeline#google_oidc EventarcPipeline#google_oidc}
	GoogleOidc *EventarcPipelineDestinationsAuthenticationConfigGoogleOidc `field:"optional" json:"googleOidc" yaml:"googleOidc"`
	// oauth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/eventarc_pipeline#oauth_token EventarcPipeline#oauth_token}
	OauthToken *EventarcPipelineDestinationsAuthenticationConfigOauthToken `field:"optional" json:"oauthToken" yaml:"oauthToken"`
}

