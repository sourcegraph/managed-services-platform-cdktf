package r2bucketlifecycle

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type R2BucketLifecycleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#account_id R2BucketLifecycle#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#bucket_name R2BucketLifecycle#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Jurisdiction of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#jurisdiction R2BucketLifecycle#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#rules R2BucketLifecycle#rules}.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

