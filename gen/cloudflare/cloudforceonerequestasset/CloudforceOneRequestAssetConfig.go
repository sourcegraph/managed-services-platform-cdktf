package cloudforceonerequestasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudforceOneRequestAssetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/cloudforce_one_request_asset#account_id CloudforceOneRequestAsset#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Page number of results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/cloudforce_one_request_asset#page CloudforceOneRequestAsset#page}
	Page *float64 `field:"required" json:"page" yaml:"page"`
	// Number of results per page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/cloudforce_one_request_asset#per_page CloudforceOneRequestAsset#per_page}
	PerPage *float64 `field:"required" json:"perPage" yaml:"perPage"`
	// UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/cloudforce_one_request_asset#request_id CloudforceOneRequestAsset#request_id}
	RequestId *string `field:"required" json:"requestId" yaml:"requestId"`
	// Asset file to upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/cloudforce_one_request_asset#source CloudforceOneRequestAsset#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

