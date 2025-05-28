package audittrailtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditTrailTokenConfig struct {
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
	// The time when the audit trail token will expire. This must be a valid ISO8601 timestamp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/audit_trail_token#expired_at AuditTrailToken#expired_at}
	ExpiredAt *string `field:"optional" json:"expiredAt" yaml:"expiredAt"`
	// When set to true will force the audit trail token to be recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/audit_trail_token#force_regenerate AuditTrailToken#force_regenerate}
	ForceRegenerate interface{} `field:"optional" json:"forceRegenerate" yaml:"forceRegenerate"`
	// Name of the organization. If omitted, organization must be defined in the provider config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/audit_trail_token#organization AuditTrailToken#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

