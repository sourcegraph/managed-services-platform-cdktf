package workersforplatformsdispatchnamespace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersForPlatformsDispatchNamespaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_for_platforms_dispatch_namespace#account_id WorkersForPlatformsDispatchNamespace#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the dispatch namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_for_platforms_dispatch_namespace#name WorkersForPlatformsDispatchNamespace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

