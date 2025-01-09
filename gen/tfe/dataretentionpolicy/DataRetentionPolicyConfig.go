package dataretentionpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataRetentionPolicyConfig struct {
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
	// delete_older_than block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/data_retention_policy#delete_older_than DataRetentionPolicy#delete_older_than}
	DeleteOlderThan *DataRetentionPolicyDeleteOlderThan `field:"optional" json:"deleteOlderThan" yaml:"deleteOlderThan"`
	// dont_delete block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/data_retention_policy#dont_delete DataRetentionPolicy#dont_delete}
	DontDelete *DataRetentionPolicyDontDelete `field:"optional" json:"dontDelete" yaml:"dontDelete"`
	// Name of the organization. If omitted, organization must be defined in the provider config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/data_retention_policy#organization DataRetentionPolicy#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// ID of the workspace that the data retention policy should apply to.
	//
	// If omitted, the data retention policy will apply to the entire organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/data_retention_policy#workspace_id DataRetentionPolicy#workspace_id}
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

