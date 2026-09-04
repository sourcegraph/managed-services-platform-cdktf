package datacloudflarespectrumapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareSpectrumApplicationConfig struct {
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
	// Zone identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/spectrum_application#zone_id DataCloudflareSpectrumApplication#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// App identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/spectrum_application#app_id DataCloudflareSpectrumApplication#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/spectrum_application#filter DataCloudflareSpectrumApplication#filter}.
	Filter *DataCloudflareSpectrumApplicationFilter `field:"optional" json:"filter" yaml:"filter"`
}

