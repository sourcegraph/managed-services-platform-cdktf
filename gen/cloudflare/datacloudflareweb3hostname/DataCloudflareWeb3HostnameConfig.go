package datacloudflareweb3hostname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareWeb3HostnameConfig struct {
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
	// Specify the identifier of the hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/web3_hostname#zone_id DataCloudflareWeb3Hostname#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Specify the identifier of the hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/web3_hostname#identifier DataCloudflareWeb3Hostname#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

