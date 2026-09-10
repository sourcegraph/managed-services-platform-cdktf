package tokenvalidationconfig


type TokenValidationConfigCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#keys TokenValidationConfig#keys}.
	Keys interface{} `field:"required" json:"keys" yaml:"keys"`
}

