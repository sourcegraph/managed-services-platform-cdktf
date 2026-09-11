package datacloudflareauthenticatedoriginpullshostnamecertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAuthenticatedOriginPullsHostnameCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/authenticated_origin_pulls_hostname_certificate#certificate_id DataCloudflareAuthenticatedOriginPullsHostnameCertificate#certificate_id}
	CertificateId *string `field:"required" json:"certificateId" yaml:"certificateId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/authenticated_origin_pulls_hostname_certificate#zone_id DataCloudflareAuthenticatedOriginPullsHostnameCertificate#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

