package datacloudflareturnstilewidgets

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareTurnstileWidgetsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/turnstile_widgets#account_id DataCloudflareTurnstileWidgets#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Direction to order widgets. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/turnstile_widgets#direction DataCloudflareTurnstileWidgets#direction}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/turnstile_widgets#filter DataCloudflareTurnstileWidgets#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/turnstile_widgets#max_items DataCloudflareTurnstileWidgets#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to order widgets by. Available values: "id", "sitekey", "name", "created_on", "modified_on".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/turnstile_widgets#order DataCloudflareTurnstileWidgets#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

