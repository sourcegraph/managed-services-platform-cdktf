package googlestoragebucketaccesscontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageBucketAccessControlConfig struct {
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
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_access_control#bucket GoogleStorageBucketAccessControl#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The entity holding the permission, in one of the following forms:   user-userId   user-email   group-groupId   group-email   domain-domain   project-team-projectId   allUsers   allAuthenticatedUsers Examples:   The user liz@example.com would be user-liz@example.com.   The group example@googlegroups.com would be   group-example@googlegroups.com.   To refer to all members of the Google Apps for Business domain   example.com, the entity would be domain-example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_access_control#entity GoogleStorageBucketAccessControl#entity}
	Entity *string `field:"required" json:"entity" yaml:"entity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_access_control#id GoogleStorageBucketAccessControl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_access_control#role GoogleStorageBucketAccessControl#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_access_control#timeouts GoogleStorageBucketAccessControl#timeouts}
	Timeouts *GoogleStorageBucketAccessControlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

