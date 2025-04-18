package accessapplication


type AccessApplicationScimConfigAuthentication struct {
	// The authentication scheme to use when making SCIM requests to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#scheme AccessApplication#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
	// URL used to generate the auth code used during token generation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#authorization_url AccessApplication#authorization_url}
	AuthorizationUrl *string `field:"optional" json:"authorizationUrl" yaml:"authorizationUrl"`
	// Client ID used to authenticate when generating a token for authenticating with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#client_id AccessApplication#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Secret used to authenticate when generating a token for authenticating with the remove SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#client_secret AccessApplication#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#password AccessApplication#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The authorization scopes to request when generating the token used to authenticate with the remove SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#scopes AccessApplication#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Token used to authenticate with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#token AccessApplication#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// URL used to generate the token used to authenticate with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#token_url AccessApplication#token_url}
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// User name used to authenticate with the remote SCIM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#user AccessApplication#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

