package googleprivilegedaccessmanagerentitlement

type GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privileged_access_manager_entitlement#steps GooglePrivilegedAccessManagerEntitlement#steps}
	Steps any `field:"required" json:"steps" yaml:"steps"`
	// Optional. Do the approvers need to provide a justification for their actions?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privileged_access_manager_entitlement#require_approver_justification GooglePrivilegedAccessManagerEntitlement#require_approver_justification}
	RequireApproverJustification any `field:"optional" json:"requireApproverJustification" yaml:"requireApproverJustification"`
}
