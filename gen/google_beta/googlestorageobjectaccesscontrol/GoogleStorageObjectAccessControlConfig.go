package googlestorageobjectaccesscontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageObjectAccessControlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_object_access_control#bucket GoogleStorageObjectAccessControl#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The entity holding the permission, in one of the following forms:   * user-{{userId}}   * user-{{email}} (such as "user-liz@example.com")   * group-{{groupId}}   * group-{{email}} (such as "group-example@googlegroups.com")   * domain-{{domain}} (such as "domain-example.com")   * project-team-{{projectId}}   * allUsers   * allAuthenticatedUsers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_object_access_control#entity GoogleStorageObjectAccessControl#entity}
	Entity *string `field:"required" json:"entity" yaml:"entity"`
	// The name of the object to apply the access control to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_object_access_control#object GoogleStorageObjectAccessControl#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// The access permission for the entity. Possible values: ["OWNER", "READER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_object_access_control#role GoogleStorageObjectAccessControl#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_object_access_control#id GoogleStorageObjectAccessControl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_object_access_control#timeouts GoogleStorageObjectAccessControl#timeouts}
	Timeouts *GoogleStorageObjectAccessControlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

