package ratelimit


type RateLimitMatchHeaders struct {
	// The name of the response header to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/rate_limit#name RateLimit#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal". Available values: "eq", "ne".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/rate_limit#op RateLimit#op}
	Op *string `field:"optional" json:"op" yaml:"op"`
	// The value of the response header, which must match exactly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/rate_limit#value RateLimit#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

