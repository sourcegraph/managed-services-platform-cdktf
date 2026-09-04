package datacloudflareturnstilewidget


type DataCloudflareTurnstileWidgetFilter struct {
	// Direction to order widgets. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/turnstile_widget#direction DataCloudflareTurnstileWidget#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter widgets by field using case-insensitive substring matching. Format: `field:value`.
	//
	// Supported fields:
	// - `name` - Filter by widget name (e.g., `filter=name:login-form`)
	// - `sitekey` - Filter by sitekey (e.g., `filter=sitekey:0x4AAA`)
	//
	// Returns 400 Bad Request if the field is unsupported or format is invalid.
	// An empty filter value returns all results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/turnstile_widget#filter DataCloudflareTurnstileWidget#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Field to order widgets by. Available values: "id", "sitekey", "name", "created_on", "modified_on".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/turnstile_widget#order DataCloudflareTurnstileWidget#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

