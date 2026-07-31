package datacloudflareleakedcredentialcheckrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareLeakedCredentialCheckRuleConfig struct {
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
	// Defines the unique ID for this custom detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/leaked_credential_check_rule#detection_id DataCloudflareLeakedCredentialCheckRule#detection_id}
	DetectionId *string `field:"required" json:"detectionId" yaml:"detectionId"`
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/leaked_credential_check_rule#zone_id DataCloudflareLeakedCredentialCheckRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

