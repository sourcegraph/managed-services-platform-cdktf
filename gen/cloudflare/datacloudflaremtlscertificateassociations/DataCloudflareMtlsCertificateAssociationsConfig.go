package datacloudflaremtlscertificateassociations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareMtlsCertificateAssociationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/mtls_certificate_associations#mtls_certificate_id DataCloudflareMtlsCertificateAssociations#mtls_certificate_id}
	MtlsCertificateId *string `field:"required" json:"mtlsCertificateId" yaml:"mtlsCertificateId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/mtls_certificate_associations#account_id DataCloudflareMtlsCertificateAssociations#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

