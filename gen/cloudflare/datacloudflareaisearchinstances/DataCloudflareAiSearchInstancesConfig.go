package datacloudflareaisearchinstances

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAiSearchInstancesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instances#account_id DataCloudflareAiSearchInstances#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instances#max_items DataCloudflareAiSearchInstances#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instances#namespace DataCloudflareAiSearchInstances#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Order By Column Name Available values: "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instances#order_by DataCloudflareAiSearchInstances#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Order By Direction Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instances#order_by_direction DataCloudflareAiSearchInstances#order_by_direction}
	OrderByDirection *string `field:"optional" json:"orderByDirection" yaml:"orderByDirection"`
	// Search by id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instances#search DataCloudflareAiSearchInstances#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

