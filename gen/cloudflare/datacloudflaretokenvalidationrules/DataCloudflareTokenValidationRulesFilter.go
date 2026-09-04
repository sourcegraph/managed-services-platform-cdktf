package datacloudflaretokenvalidationrules


type DataCloudflareTokenValidationRulesFilter struct {
	// Action to take on requests that match operations included in `selector` and fail `expression`. Available values: "log", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules#action DataCloudflareTokenValidationRules#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Toggle rule on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules#enabled DataCloudflareTokenValidationRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Select rules with this host in `include`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules#host DataCloudflareTokenValidationRules#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Select rules with this host in `include`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules#hostname DataCloudflareTokenValidationRules#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Select rules with these IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules#id DataCloudflareTokenValidationRules#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Select rules using any of these token configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules#token_configuration DataCloudflareTokenValidationRules#token_configuration}
	TokenConfiguration *[]*string `field:"optional" json:"tokenConfiguration" yaml:"tokenConfiguration"`
}

