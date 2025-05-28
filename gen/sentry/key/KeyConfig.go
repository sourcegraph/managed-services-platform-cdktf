package key

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyConfig struct {
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
	// The name of the client key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/key#name Key#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/key#organization Key#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The project of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/key#project Key#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The JavaScript loader script configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/key#javascript_loader_script Key#javascript_loader_script}
	JavascriptLoaderScript *KeyJavascriptLoaderScript `field:"optional" json:"javascriptLoaderScript" yaml:"javascriptLoaderScript"`
	// Number of events that can be reported within the rate limit window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/key#rate_limit_count Key#rate_limit_count}
	RateLimitCount *float64 `field:"optional" json:"rateLimitCount" yaml:"rateLimitCount"`
	// Length of time in seconds that will be considered when checking the rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/key#rate_limit_window Key#rate_limit_window}
	RateLimitWindow *float64 `field:"optional" json:"rateLimitWindow" yaml:"rateLimitWindow"`
}

