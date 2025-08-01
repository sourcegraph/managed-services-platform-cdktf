package datacloudflarecloudforceonerequestasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCloudforceOneRequestAssetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/cloudforce_one_request_asset#account_id DataCloudflareCloudforceOneRequestAsset#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/cloudforce_one_request_asset#asset_id DataCloudflareCloudforceOneRequestAsset#asset_id}
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/cloudforce_one_request_asset#request_id DataCloudflareCloudforceOneRequestAsset#request_id}
	RequestId *string `field:"required" json:"requestId" yaml:"requestId"`
}

