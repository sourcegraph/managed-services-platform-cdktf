package byoipprefix

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ByoIpPrefixConfig struct {
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
	// Autonomous System Number (ASN) the prefix will be advertised under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/byo_ip_prefix#asn ByoIpPrefix#asn}
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// IP Prefix in Classless Inter-Domain Routing format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/byo_ip_prefix#cidr ByoIpPrefix#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Identifier of a Cloudflare account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/byo_ip_prefix#account_id ByoIpPrefix#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Whether Cloudflare is allowed to generate the LOA document on behalf of the prefix owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/byo_ip_prefix#delegate_loa_creation ByoIpPrefix#delegate_loa_creation}
	DelegateLoaCreation interface{} `field:"optional" json:"delegateLoaCreation" yaml:"delegateLoaCreation"`
	// Description of the prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/byo_ip_prefix#description ByoIpPrefix#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Identifier for the uploaded LOA document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/byo_ip_prefix#loa_document_id ByoIpPrefix#loa_document_id}
	LoaDocumentId *string `field:"optional" json:"loaDocumentId" yaml:"loaDocumentId"`
}

