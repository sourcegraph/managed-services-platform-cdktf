package datacloudflareaccountpermissiongroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAccountPermissionGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_permission_group#account_id DataCloudflareAccountPermissionGroup#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Permission Group identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_permission_group#permission_group_id DataCloudflareAccountPermissionGroup#permission_group_id}
	PermissionGroupId *string `field:"required" json:"permissionGroupId" yaml:"permissionGroupId"`
}

