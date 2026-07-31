package customssl


type CustomSslGeoRestrictions struct {
	// Available values: "us", "eu", "highest_security".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/custom_ssl#label CustomSsl#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

