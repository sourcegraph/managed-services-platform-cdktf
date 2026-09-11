package zerotrustgatewaycertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_certificate#account_id ZeroTrustGatewayCertificate#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Whether to activate the certificate on Cloudflare's edge.
	//
	// When true, the certificate will be activated. When false, the certificate will be deactivated at the edge. This is a Terraform-only field and does not appear in the API response. Monitor `binding_status` for the activation status. Once a certificate is activated, you may use the certificate to intercept traffic
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_certificate#activate ZeroTrustGatewayCertificate#activate}
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// Sets the certificate validity period in days (range: 1-10,950 days / ~30 years).
	//
	// Defaults to 1,825 days (5 years). **Important**: This field is only settable during the certificate creation.  Certificates becomes immutable after creation - use the `/activate` and `/deactivate` endpoints to manage certificate lifecycle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_certificate#validity_period_days ZeroTrustGatewayCertificate#validity_period_days}
	ValidityPeriodDays *float64 `field:"optional" json:"validityPeriodDays" yaml:"validityPeriodDays"`
}

