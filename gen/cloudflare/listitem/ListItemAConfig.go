package listitem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ListItemAConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list_item#account_id ListItemA#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The list identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list_item#list_id ListItemA#list_id}
	ListId *string `field:"required" json:"listId" yaml:"listId"`
	// Autonomous system number to include in the list. Must provide only one of: `ip`, `asn`, `redirect`, `hostname`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list_item#asn ListItemA#asn}
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// An optional comment for the item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list_item#comment ListItemA#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list_item#hostname ListItemA#hostname}
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// IP address to include in the list. Must provide only one of: `ip`, `asn`, `redirect`, `hostname`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list_item#ip ListItemA#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list_item#redirect ListItemA#redirect}
	Redirect interface{} `field:"optional" json:"redirect" yaml:"redirect"`
}

