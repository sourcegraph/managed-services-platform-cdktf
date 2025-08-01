package datacloudflareaccountapitokenpermissiongroupslist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAccountApiTokenPermissionGroupsListConfig struct {
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
	// Account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_api_token_permission_groups_list#account_id DataCloudflareAccountApiTokenPermissionGroupsList#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_api_token_permission_groups_list#max_items DataCloudflareAccountApiTokenPermissionGroupsList#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Filter by the name of the permission group. The value must be URL-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_api_token_permission_groups_list#name DataCloudflareAccountApiTokenPermissionGroupsList#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Filter by the scope of the permission group. The value must be URL-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_api_token_permission_groups_list#scope DataCloudflareAccountApiTokenPermissionGroupsList#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

