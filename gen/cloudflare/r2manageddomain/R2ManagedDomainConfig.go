package r2manageddomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type R2ManagedDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_managed_domain#account_id R2ManagedDomain#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_managed_domain#bucket_name R2ManagedDomain#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Whether to enable public bucket access at the r2.dev domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_managed_domain#enabled R2ManagedDomain#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Jurisdiction of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_managed_domain#jurisdiction R2ManagedDomain#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
}

