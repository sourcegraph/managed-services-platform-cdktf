package accessapplication


type AccessApplicationSaasAppRefreshTokenOptions struct {
	// How long a refresh token will be valid for after creation.
	//
	// Valid units are `m`, `h` and `d`. Must be longer than 1m.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#lifetime AccessApplication#lifetime}
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
}

