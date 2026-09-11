package datacloudflareworkerscustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWorkersCustomDomainConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/workers_custom_domain#account_id DataCloudflareWorkersCustomDomain#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// ID of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/workers_custom_domain#domain_id DataCloudflareWorkersCustomDomain#domain_id}
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/workers_custom_domain#filter DataCloudflareWorkersCustomDomain#filter}.
	Filter *DataCloudflareWorkersCustomDomainFilter `field:"optional" json:"filter" yaml:"filter"`
}

