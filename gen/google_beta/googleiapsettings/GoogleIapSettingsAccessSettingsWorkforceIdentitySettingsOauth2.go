package googleiapsettings


type GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2 struct {
	// The OAuth 2.0 client ID registered in the workforce identity federation OAuth 2.0 Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_iap_settings#client_id GoogleIapSettings#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Input only. The OAuth 2.0 client secret created while registering the client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_iap_settings#client_secret GoogleIapSettings#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

