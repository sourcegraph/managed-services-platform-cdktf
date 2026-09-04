package workersscript


type WorkersScriptBindingsSimple struct {
	// The rate limit value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#limit WorkersScript#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The rate limit period in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#period WorkersScript#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
}

