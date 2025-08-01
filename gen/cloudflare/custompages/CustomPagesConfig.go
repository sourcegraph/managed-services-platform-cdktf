package custompages

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomPagesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/custom_pages#identifier CustomPages#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The custom page state. Available values: "default", "customized".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/custom_pages#state CustomPages#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The URL associated with the custom page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/custom_pages#url CustomPages#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/custom_pages#account_id CustomPages#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/custom_pages#zone_id CustomPages#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

