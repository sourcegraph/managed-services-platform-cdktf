package datacloudflaresnippet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareSnippetConfig struct {
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
	// Identify the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/snippet#snippet_name DataCloudflareSnippet#snippet_name}
	SnippetName *string `field:"required" json:"snippetName" yaml:"snippetName"`
	// Use this field to specify the unique ID of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/snippet#zone_id DataCloudflareSnippet#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

