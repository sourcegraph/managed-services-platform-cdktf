package zerotrustaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessPolicyConfig struct {
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
	// The action Access will take if a user matches this policy.
	//
	// Infrastructure application policies can only use the Allow action.
	// Available values: "allow", "deny", "non_identity", "bypass".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#decision ZeroTrustAccessPolicy#decision}
	Decision *string `field:"required" json:"decision" yaml:"decision"`
	// The name of the Access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#name ZeroTrustAccessPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#account_id ZeroTrustAccessPolicy#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Administrators who can approve a temporary authentication request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#approval_groups ZeroTrustAccessPolicy#approval_groups}
	ApprovalGroups interface{} `field:"optional" json:"approvalGroups" yaml:"approvalGroups"`
	// Requires the user to request access from an administrator at the start of each session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#approval_required ZeroTrustAccessPolicy#approval_required}
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
	// The rules that define how users may connect to targets secured by your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#connection_rules ZeroTrustAccessPolicy#connection_rules}
	ConnectionRules *ZeroTrustAccessPolicyConnectionRules `field:"optional" json:"connectionRules" yaml:"connectionRules"`
	// Rules evaluated with a NOT logical operator.
	//
	// To match the policy, a user cannot meet any of the Exclude rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#exclude ZeroTrustAccessPolicy#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#include ZeroTrustAccessPolicy#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// Require this application to be served in an isolated browser for users matching this policy.
	//
	// 'Client Web Isolation' must be on for the account in order to use this feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#isolation_required ZeroTrustAccessPolicy#isolation_required}
	IsolationRequired interface{} `field:"optional" json:"isolationRequired" yaml:"isolationRequired"`
	// Configures multi-factor authentication (MFA) settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#mfa_config ZeroTrustAccessPolicy#mfa_config}
	MfaConfig *ZeroTrustAccessPolicyMfaConfig `field:"optional" json:"mfaConfig" yaml:"mfaConfig"`
	// A custom message that will appear on the purpose justification screen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#purpose_justification_prompt ZeroTrustAccessPolicy#purpose_justification_prompt}
	PurposeJustificationPrompt *string `field:"optional" json:"purposeJustificationPrompt" yaml:"purposeJustificationPrompt"`
	// Require users to enter a justification when they log in to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#purpose_justification_required ZeroTrustAccessPolicy#purpose_justification_required}
	PurposeJustificationRequired interface{} `field:"optional" json:"purposeJustificationRequired" yaml:"purposeJustificationRequired"`
	// Rules evaluated with an AND logical operator.
	//
	// To match the policy, a user must meet all of the Require rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#require ZeroTrustAccessPolicy#require}
	Require interface{} `field:"optional" json:"require" yaml:"require"`
	// The amount of time that tokens issued for the application will be valid.
	//
	// Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#session_duration ZeroTrustAccessPolicy#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
}

