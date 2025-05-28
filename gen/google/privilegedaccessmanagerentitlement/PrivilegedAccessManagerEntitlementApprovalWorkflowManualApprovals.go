package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/privileged_access_manager_entitlement#steps PrivilegedAccessManagerEntitlement#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// Optional. Do the approvers need to provide a justification for their actions?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/privileged_access_manager_entitlement#require_approver_justification PrivilegedAccessManagerEntitlement#require_approver_justification}
	RequireApproverJustification interface{} `field:"optional" json:"requireApproverJustification" yaml:"requireApproverJustification"`
}

