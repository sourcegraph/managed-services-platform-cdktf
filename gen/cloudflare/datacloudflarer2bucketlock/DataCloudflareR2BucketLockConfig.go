package datacloudflarer2bucketlock

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareR2BucketLockConfig struct {
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
	// Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/r2_bucket_lock#account_id DataCloudflareR2BucketLock#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/r2_bucket_lock#bucket_name DataCloudflareR2BucketLock#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

