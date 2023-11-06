package googlecomputebackendservice


type GoogleComputeBackendServiceIap struct {
	// OAuth2 Client ID for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_backend_service#oauth2_client_id GoogleComputeBackendService#oauth2_client_id}
	Oauth2ClientId *string `field:"required" json:"oauth2ClientId" yaml:"oauth2ClientId"`
	// OAuth2 Client Secret for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_backend_service#oauth2_client_secret GoogleComputeBackendService#oauth2_client_secret}
	Oauth2ClientSecret *string `field:"required" json:"oauth2ClientSecret" yaml:"oauth2ClientSecret"`
}

