package datacloudflarecertificateauthoritieshostnameassociations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCertificateAuthoritiesHostnameAssociationsConfig struct {
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
	// The UUID to match against for a certificate that was uploaded to the mTLS Certificate Management endpoint.
	//
	// If no mtls_certificate_id is given, the results will be the hostnames associated to your active Cloudflare Managed CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/certificate_authorities_hostname_associations#mtls_certificate_id DataCloudflareCertificateAuthoritiesHostnameAssociations#mtls_certificate_id}
	MtlsCertificateId *string `field:"optional" json:"mtlsCertificateId" yaml:"mtlsCertificateId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/certificate_authorities_hostname_associations#zone_id DataCloudflareCertificateAuthoritiesHostnameAssociations#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

