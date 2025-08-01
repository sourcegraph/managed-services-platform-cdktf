package datagooglestoragebucketobjectcontent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleStorageBucketObjectContentConfig struct {
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
	// The name of the containing bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_storage_bucket_object_content#bucket DataGoogleStorageBucketObjectContent#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_storage_bucket_object_content#name DataGoogleStorageBucketObjectContent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Data as string to be uploaded.
	//
	// Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_storage_bucket_object_content#content DataGoogleStorageBucketObjectContent#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_storage_bucket_object_content#id DataGoogleStorageBucketObjectContent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

