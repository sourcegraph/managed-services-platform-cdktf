package tokenvalidationconfig


type TokenValidationConfigCredentialsKeys struct {
	// Algorithm Available values: "RS256", "RS384", "RS512", "PS256", "PS384", "PS512", "ES256", "ES384".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#alg TokenValidationConfig#alg}
	Alg *string `field:"required" json:"alg" yaml:"alg"`
	// Key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#kid TokenValidationConfig#kid}
	Kid *string `field:"required" json:"kid" yaml:"kid"`
	// Key Type Available values: "RSA", "EC".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#kty TokenValidationConfig#kty}
	Kty *string `field:"required" json:"kty" yaml:"kty"`
	// Curve Available values: "P-256", "P-384".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#crv TokenValidationConfig#crv}
	Crv *string `field:"optional" json:"crv" yaml:"crv"`
	// RSA exponent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#e TokenValidationConfig#e}
	E *string `field:"optional" json:"e" yaml:"e"`
	// RSA modulus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#n TokenValidationConfig#n}
	N *string `field:"optional" json:"n" yaml:"n"`
	// X EC coordinate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#x TokenValidationConfig#x}
	X *string `field:"optional" json:"x" yaml:"x"`
	// Y EC coordinate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_config#y TokenValidationConfig#y}
	Y *string `field:"optional" json:"y" yaml:"y"`
}

