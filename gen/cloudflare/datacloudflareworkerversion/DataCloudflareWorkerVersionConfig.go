package datacloudflareworkerversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWorkerVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/worker_version#account_id DataCloudflareWorkerVersion#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Identifier for the version, which can be ID or the literal "latest" to operate on the most recently created version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/worker_version#version_id DataCloudflareWorkerVersion#version_id}
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
	// Identifier for the Worker, which can be ID or name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/worker_version#worker_id DataCloudflareWorkerVersion#worker_id}
	WorkerId *string `field:"required" json:"workerId" yaml:"workerId"`
	// Whether to include the `modules` property of the version in the response, which contains code and sourcemap content and may add several megabytes to the response size.
	//
	// Available values: "modules".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/worker_version#include DataCloudflareWorkerVersion#include}
	Include *string `field:"optional" json:"include" yaml:"include"`
}

