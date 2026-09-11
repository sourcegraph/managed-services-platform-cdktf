package apishield


type ApiShieldAuthIdCharacteristics struct {
	// The name of the characteristic field, i.e., the header or cookie name. When using type "jwt", this must be a claim location expressed as `$(token_config_id):$(json_path)`, where `token_config_id` is the ID of the token configuration used in validating the JWT, and `json_path` is a RFC 9535 JSONPath expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/api_shield#name ApiShield#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of characteristic. Available values: "header", "cookie", "jwt".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/api_shield#type ApiShield#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

